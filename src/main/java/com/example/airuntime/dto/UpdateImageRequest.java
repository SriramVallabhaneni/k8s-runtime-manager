package com.example.airuntime.dto;

import jakarta.validation.constraints.NotBlank;

public class UpdateImageRequest {
    @NotBlank(message = "Container image is required")
    private String image;

    public String getImage() {
        return image;
    }
}