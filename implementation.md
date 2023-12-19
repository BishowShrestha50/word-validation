# Implementation Details

This document provides an overview of the core logic and implementation details of the English Word Validation API.

## Core Logic
 - Starting of the program the api loads a list of 350,000 English words from `words.txt` file into memory.
- The `GetAllWords` function loads the list of words in the `words.txt` and creates a map for efficient lookups.
- The `WordsHandler` function checks if the provided word is present in the list of valid words. 

The English Word Validation API follows a simple architecture:



**Controller**
- Controller is responsible for core part of the word validation api. Here the request for validate api is received and the route of the api to `WordsHanlder` function is done which checks if the provided word is present in the list of valid word.
- All the api related work is done in this controller folder.
- The API exposes a single endpoint (/v1/validate) for word validation.

**Service**

- Service folder is responsible for doing the vaidation related tasks.
- They are intermediary layers between controllers and models.
- There is the interface defined through which controller can call the method in the service.
- The `GetAllWords` method loads the list of words in the words.txt and creates a map for efficient lookups.

**Model**

- The model represents the data structure of the application. It defines the basic schema, structure, and rules around the data.
- The model rquired for the validation api is initialized here.

## Deployment and Containerization

### Dockerfile

- The provided `Dockerfile` uses the official Golang base image to build the application.
- The application binary is compiled within the Docker image.
- The resulting image is tagged and ready for deployment.

### Kubernetes Deployment 

- **Deployment :**
  - Establishes a Kubernetes Deployment with a single replica.
  - Sets the CPU and memory resource constraints and allocations for the container.
  - Specifies probes for readiness and liveness to ensure health checks.

- **Service:**
  - Creates a Kubernetes Service to make the API accessible within the cluster.

- **Ingress:**
  - Sets up an Ingress resource for external entry, assuming AWS infrastructure.
  - Specifies the DNS name, namespace, and AWS ALB integration annotations.


## Conclusion

This document offers an insight into the foundational logic of the English Word Validation API, outlining its core functionalities, deployment approaches, and prospective enhancements. Its purpose is to aid developers, maintainers, and contributors in comprehending the architecture and strategic decisions of the project, facilitating a deeper understanding of its structure and design principles.
