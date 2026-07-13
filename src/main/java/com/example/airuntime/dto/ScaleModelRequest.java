package com.example.airuntime.dto;

import jakarta.validation.constraints.Min;

public class ScaleModelRequest {

    @Min(value = 1, message = "Replicas must be at least 1")
    private int replicas;

    public int getReplicas() {
        return replicas;
    }
}
