# word-validation

# English Word Validation API

This project implements a simple REST API for validating English words. It provides an endpoint that takes a word as a parameter and responds with whether the word is valid based on a pre-defined list of over 350,000 English words.

## Table of Contents

- [Setup](#setup)
  - [Prerequisites](#prerequisites)
  - [Build and Run Locally](#build-and-run-locally)
- [Docker Setup](#docker-setup)
  - [Build Docker Image](#build-docker-image)
  - [Run Container](#run-container)
- [Kubernetes Setup](#kubernetes-setup)
  - [Deploy to Kubernetes Cluster](#deploy-to-kubernetes-cluster)
  - [Access the Service](#access-the-service)
- [API Usage](#api-usage)
  - [Endpoint](#endpoint)
  - [Example](#example)

## Setup

### Prerequisites

- Go (v1.17 or later) 

### Build and Run Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/BishowShrestha50/word-validation.git
   ```

2. Go to the project directory:

   ```bash 
   cd word-validation
   ```

 3. Run the project

    ```bash
    go run main.go 
    ``` 

### Docker Setup

#### Build Docker Image

1. Build docker image with your registry username and repo name

   ```
   docker build -t <USERNAME>/<REPO-NAME>:<TAG> .
   ```
   Use sample docker image:
   ```
   bishowshrestha50/word-validation:latest
   ```

2. Push docker image to your registry:

   ```
   docker push  <USERNAME>/<REPO-NAME>:<TAG>
   ```
### Run container
    
    docker run --rm --env-file=<ENV-PATH> -p 8080:8080 <USERNAME>/<REPO-NAME>:<TAG>
    


### Kubernetes Setup

#### Deploy to Kubernetes Cluster

1. Create deployment

   ```
   kubectl apply -f kubernetes/deployment.yaml
   ```

2. Create service of the deployment

   ```
   kubectl apply -f kubernetes/service.yaml
   ```   

3. Deploy ingress controller 
    
    For AWS infrastructure (LB), leveraging aws-lb-controller
   ```
   helm install aws-load-balancer-controller eks/aws-load-balancer-controller -n kube-system --set clusterName=my-cluster --set serviceAccount.create=false --set serviceAccount.name=aws-load-balancer-controller 
   ```
3. Create ingress route for external access

   ```
   kubectl apply -f kubernetes/ingress.yaml
   ```

#### Access the  service

   ```
   kubectl port-forward svc/<service-name> -p 8080:8080
   ```

### API usage

#### Endpoint

1. Endpoint URL
   ```
   GET <BASE_URL>/v1/validate
   ```
2. Parameters
   ```
   word: string value of the word to be searched 
   ```

#### Example

Valid Response
```
Request: GET /v1/validate?word=boat
Response: { "valid": true }
```
Invalid Response

```
Request: GET /v1/validate?word=bbb
Response: { "valid": false }
```