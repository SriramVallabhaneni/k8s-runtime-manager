# AI Runtime Manager

A cloud-native Java application for deploying and managing AI runtimes on Kubernetes.

AI Runtime Manager exposes a REST API that provisions and manages AI inference runtimes (currently Ollama-based) inside Kubernetes clusters. The application automates workload deployment, scaling, image updates, restarts, persistent model storage, and Kubernetes resource management while running as a Kubernetes-native service itself.

---

## Features

- Deploy supported AI runtimes through a REST API
- Automatically create Kubernetes Deployments, Services, and PersistentVolumeClaims
- Automatically download supported AI models during initialization
- Persist downloaded models across Pod recreation using PersistentVolumeClaims
- Scale, restart, update, inspect, and delete AI runtime deployments
- Kubernetes-native deployment using Helm
- RBAC-based in-cluster authentication
- OpenAPI (Swagger) documentation
- Request validation and centralized exception handling

---

## Architecture

```text
                 Client
                    │
                    ▼
          Spring Boot REST API
          (AI Runtime Manager)
                    │
                    ▼
        Kubernetes Java Client
                    │
                    ▼
            Kubernetes API Server
                    │
     ┌──────────────┼──────────────┐
     ▼              ▼              ▼
 Deployment      Service         PVC
     │
     ▼
Init Container
(download AI model)
     │
     ▼
 Ollama Runtime
     │
     ▼
AI Inference
```

---

## Technology Stack

| Category | Technologies |
|----------|--------------|
| Language | Java 24 |
| Framework | Spring Boot |
| Containerization | Docker |
| Orchestration | Kubernetes |
| Packaging | Helm |
| API Documentation | OpenAPI / Swagger |
| Build Tool | Maven |
| AI Runtime | Ollama |
| Kubernetes Client | Kubernetes Java Client |

---

## Current Supported Models

| Model | Runtime |
|------|---------|
| TinyLlama | Ollama |
| Llama 3 | Ollama |

The model registry can be extended by adding additional supported models.

---

## REST API

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | `/models` | Deploy AI runtime |
| GET | `/models` | List deployments |
| GET | `/models/{name}` | Deployment status |
| PATCH | `/models/{name}/scale` | Scale deployment |
| PATCH | `/models/{name}/image` | Update container image |
| POST | `/models/{name}/restart` | Restart deployment |
| DELETE | `/models/{name}` | Delete deployment |

Interactive documentation is available through Swagger:

```
/swagger-ui/index.html
```

---

## Example Deployment Request

```json
{
    "deploymentName": "tinyllama-runtime",
    "model": "tinyllama",
    "replicas": 1
}
```

The runtime manager automatically determines:

- Container image
- Container port
- Persistent storage
- Kubernetes resources

The client only specifies the AI model to deploy.

---

## Running Locally

Clone the repository:

```bash
git clone https://github.com/<username>/k8s-runtime-manager.git
cd k8s-runtime-manager
```

Build:

```bash
mvn clean package
```

Run:

```bash
mvn spring-boot:run
```

---

## Running on Kubernetes

Build the Docker image:

```bash
docker build -t k8s-runtime-manager:1.0 .
```

Deploy with Helm:

```bash
helm install runtime-manager ./helm/runtime-manager
```

---

## Project Structure

```
src/
 ├── controller/
 ├── dto/
 ├── service/
 ├── model/
 └── config/

helm/
k8s/
```

---

## Future Improvements

- Kubernetes Custom Resource Definitions (CRDs)
- Go-based Kubernetes Operator
- Multi-model runtime support
- GPU-aware scheduling
- Multi-replica AI runtime support
- Additional AI runtime backends

---

## Why this project?

The goal of this project is to explore Kubernetes-native application development by building a backend service that automates the deployment and lifecycle management of AI inference runtimes.

Rather than exposing raw Kubernetes resources to clients, the application provides a higher-level API centered around AI models while using Kubernetes as the underlying orchestration platform.