package com.example.airuntime.config;

import io.kubernetes.client.openapi.ApiClient;
import io.kubernetes.client.util.Config;
import io.kubernetes.client.openapi.Configuration;
import org.springframework.context.annotation.Bean;

@org.springframework.context.annotation.Configuration
public class KubernetesConfig {

    @Bean
    public ApiClient apiClient() throws Exception {
        ApiClient client;

        if (System.getenv("KUBERNETES_SERVICE_HOST") != null) {
            client = Config.fromCluster();
        } else {
            client = Config.defaultClient();
        }

        Configuration.setDefaultApiClient(client);
        return client;
    }
}
