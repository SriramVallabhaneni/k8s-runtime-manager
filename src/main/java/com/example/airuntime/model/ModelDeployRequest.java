package com.example.airuntime.model;

public class ModelDeployRequest {
    private String name;
    private String image;
    private int replicas;
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
