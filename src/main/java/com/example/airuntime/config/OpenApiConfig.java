package com.example.airuntime.config;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Contact;
import io.swagger.v3.oas.models.info.Info;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class OpenApiConfig {

    @Bean
    public OpenAPI runtimeManagerOpenAPI() {

        return new OpenAPI()
                .info(
                        new Info()
                                .title("K8s Runtime Manager API")
                                .description("""
                                        REST API for managing Kubernetes workloads.

                                        Features:
                                        • Deploy workloads
                                        • Scale deployments
                                        • Update container images
                                        • Restart deployments
                                        • Delete deployments
                                        """)
                                .version("1.0.0")
                                .contact(
                                        new Contact()
                                                .name("Sriram Vallabhaneni")
                                                .url("https://github.com/SriramVallabhaneni/k8s-runtime-manager")
                                )
                );
    }


}
