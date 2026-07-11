package com.example.airuntime.exception;

import com.example.airuntime.dto.ErrorResponse;
import io.kubernetes.client.openapi.ApiException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import java.util.List;

@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<ErrorResponse> handleValidationErrors(MethodArgumentNotValidException exception) {
        List<String> messages = exception.getBindingResult()
                .getFieldErrors()
                .stream()
                .map(error -> error.getField() + ": " + error.getDefaultMessage())
                .toList();

        ErrorResponse response = new ErrorResponse(
                400,
                "Bad Request",
                messages
        );

        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(response);
    }

    @ExceptionHandler(ApiException.class)
    public ResponseEntity<ErrorResponse> handleKubernetesErrors(ApiException exception) {
        ErrorResponse response = new ErrorResponse(
                exception.getCode(),
                "Kubernetes API Error",
                List.of(exception.getResponseBody() == null ? exception.getMessage() : exception.getResponseBody())
        );

        return ResponseEntity.status(exception.getCode()).body(response);
    }

    @ExceptionHandler(Exception.class)
    public ResponseEntity<ErrorResponse> handleGeneralErrors(Exception exception) {
        ErrorResponse response = new ErrorResponse(
                500,
                "Internal Server Error",
                List.of(exception.getMessage())
        );

        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR).body(response);
    }
}