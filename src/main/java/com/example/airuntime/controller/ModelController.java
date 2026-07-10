package com.example.airuntime.controller;

import com.example.airuntime.model.ModelDeployRequest;
import com.example.airuntime.service.KubernetesDeploymentService;
import io.kubernetes.client.openapi.models.V1Deployment;
import io.kubernetes.client.openapi.models.V1DeploymentList;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/models")
public class ModelController {

    private final KubernetesDeploymentService service;

    public ModelController(KubernetesDeploymentService service) {
        this.service = service;
    }

    @PostMapping
    public String deployModel(@RequestBody ModelDeployRequest request) throws Exception {
        return service.deployModel(request);
    }

    @GetMapping
    public V1DeploymentList listModels() throws Exception {
        return service.listModels();
    }

    @GetMapping("/{name}")
    public V1Deployment getModel(@PathVariable String name) throws Exception {
        return service.getModel(name);
    }

    @DeleteMapping("/{name}")
    public String deleteModel(@PathVariable String name) throws Exception {
        return service.deleteModel(name);
    }
}
