package com.example.airuntime.dto;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Pattern;


public class ModelDeployRequest {

    @NotBlank(message = "Model name is required")
    @Pattern(
            regexp = "^[a-z0-9]([-a-z0-9]*[a-z0-9])?$",
            message = "Model name must be lowercase and Kubernetes-compatible")
    private String name;

    @NotBlank(message = "Container image is required")
    private String image;

    @Min(value = 1, message = "Replicas must be at least 1")
    private int replicas;

    @Min(value = 1, message = "Port must be at least 1")
    @Max(value = 65535, message = "Port must be at most 65535")
    private int port;

    public String getName() {
        return name;
    }

    public String getImage() {
        return image;
    }

    public int getReplicas() {
        return replicas;
    }

    public int getPort() {
        return port;
    }
}
