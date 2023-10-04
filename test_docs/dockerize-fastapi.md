# How to Dockerize and Deploy a Fast API Application to Kubernetes Cluster

![img](https://sweetcode.io/wp-content/uploads/2022/11/Bravin-Wasike-60x60.png)

BY [BRAVIN WASIKE](https://sweetcode.io/author/bwasike/)

[Docker](https://www.docker.com/) is one of the most popular containerization tools that leverage OS virtualization. It creates Docker Containers that run in an isolated environment from the original application. Docker containers have all the configuration and files to run an application. In this tutorial, you will learn how to Dockerize and Deploy a [Fast API](https://fastapi.tiangolo.com/) application to Kubernetes. 

First, we will create and run a Fast API application locally. We will then Dockerize the Fast API application and run it as a Docker container. Finally, we will deploy the Dockerized Fast API application to Kubernetes.

[Kubernetes](https://kubernetes.io/) is an open-source platform that automatically manages the deployment of Docker containers. Kubernetes orchestrates application containers and hosts them in a Kubernetes cluster. A Kubernetes cluster has nodes (master and worker nodes) that run and manage the hosted/deployed container applications. The deployed containers run in a Kubernetes environment known as pods. A pod is the smallest deployable unit that runs in a Kubernetes cluster. 

In this tutorial, we will deploy the Dockerized Fast API application to a local Kubernetes Cluster known as [Minikube](https://minikube.sigs.k8s.io/docs/start/). We will talk more about Minikube later. Let’s start working on our Fast API application project!

##### Prerequisites

To easily follow along with this tutorial and deploy your application to Kubernetes, you need to understand [Docker](https://docs.docker.com/get-docker/). 

Before deploying your application, ensure you have the following on your computer:

- [Python ≥v3.7](https://www.python.org/)
- [Python Pip](https://pypi.org/project/pip/) package manager
- [VS Code](https://code.visualstudio.com/) code editor
- [Docker Desktop](https://www.docker.com/products/docker-desktop/)

# Getting Started With Fast API

FastAPI is a popular Python Web framework for creating RESTful APIs. It has Pydantic and Python-type hints for data serialization, deserialization, and validation. In this tutorial, we will use FastAPI to create a simple “Hello World” application. We will test and run the application locally. 

FastAPI requires a [ASGI server](https://asgi.readthedocs.io/en/latest/) to run the application in production such as [Uvicorn](https://www.uvicorn.org/). Let’s install the Fast API framework and Uvicorn using the following ‘pip’ commands:

`pip install fastapi`

`pip install "uvicorn[standard]"`

After the installation process, create a new folder named `fast-api` and open it with VS Code. In the folder, create a new ‘main.py’ file and add the following Python code:

```python
from typing import Union

from fastapi import FastAPI

app = FastAPI()

@app.get("/")

def read_root():
  return {"Hello": "World"}
```



The Python snippet above will create a ‘route’ that will return “Hello”: “World”.

To run the application, ‘cd’ into the ‘fast-api’ folder and execute this ‘uvicorn’ command in your terminal:

uvicorn main:app

The command will start your application on http://127.0.0.1:8000 as shown below:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_1.png)

You can then type the URL above in your web browser to view the application:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_2.png)

Our FastAPI application is running, the next step is to Dockerize the FastAPI application.

## Dockerizing the FastAPI Application

We will run the FastAPI application inside a Docker Container. To Dockerize the FastAPI application, we need to follow these steps:

1. Create a ‘requirements.txt’ file.
2. Write a Dockerfile for the FastAPI application.
3. Build a Docker Image for the FastAPI application using the Dockerfile.
4. Launch a Docker Container using the Docker Image.

Let’s apply these steps.

##### Create a ‘requirements.txt’ File

A ‘requirements.txt’ file contains the framework and the dependencies for the FastAPI application. Docker will install all these requirements when creating a Docker image for the FastAPI application.

In the ‘fast-api’ folder, create a new file named ‘requirements.txt’. Open the file and add the following:

```
fastapi
uvicorn
```



##### Write a Dockerfile for the FastAPI Application

Dockerfile is an executable file that contains a list of commands or instructions for creating a Docker Image. Docker Engine will execute all these commands and build a Docker image. In the ‘fast-api’ folder, create a new file named ‘requirements.txt’. 

Open the file and add the following:

```dockerfile
#It instructs Docker Engine to use official python:3.10 as the base image
FROM python:3.10
#It creates a working directory(app) for the Docker image and container
WORKDIR /app
#It copies the framework and the dependencies for the FastAPI application into the working directory
COPY requirements.txt .
#It will install the framework and the dependencies in the `requirements.txt` file.
RUN pip install -r requirements.txt
#It will copy the remaining files and the source code from the host `fast-api` folder to the `app` container working directory
COPY . .
#It will expose the FastAPI application on port `8000` inside the container
EXPOSE 8000
#It is the command that will start and run the FastAPI application container
CMD ["uvicorn", "main:app", "--host", "0.0.0.0"]
```





## Build a Docker Image for the FastAPI Application Using the Dockerfile

We will build the Docker Image for the FastAPI application using the instructions in the Dockerfile. You will ‘cd’ into the ‘fast-api’ folder in your terminal and run this ‘docker’ command:

docker build -t bravinwasike/fast-api .

**NOTE:** Ensure you name the Docker image using your Docker Hub account user name. It will be easy to push the Docker Image to the Docker Hub repository. The command will display the following output in your terminal:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_3.png)

## Launch the Fast API Application Docker Container Using the Docker Image

We will launch the Fast API application Docker Container using the Docker Image. When you run a Docker Image, it will launch a Docker Container for the application. In your terminal, run this ‘docker’ command:

“`docker

docker run -p 8000:8000 bravinwasike/fast-api

“`

The command will start your application on http://127.0.0.1:8000 as shown below:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_4.png)

You can then type the URL above in your web browser to view the application:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_5-1024x418.png)

We have Dockerized the Fast API application. Let’s push the Docker Image to Docker Hub.

## Pushing the Docker Image for the Fast API Application to Docker Hub

To push the Docker Image for the Fast API application to Docker Hub, you need to login into your Docker Hub account from your terminal:

`docker login`

You will then push the Docker image using the following ‘docker’ command:

`docker push bravinwasike/fast-api`

You must push the Docker Image that you have created to Docker Hub. We will use this Docker Image when deploying the Fast API application to Minikube. Let’s get started with Minikube.

## Getting Started With Minikube

Minikube is a Kubernetes engine that allows DevOps engineers to create a Kubernetes cluster on their local machines. The Minikube Kubernetes Cluster will only have a single node for managing the deployed application. Minikube is a simple version of a cloud-based Kubernetes Cluster with multiple masters and worker nodes. We will use Minikube to host our Dockerized FastAPI application.

Minikube is easy to set up because it usually comes with a Docker installation. As long as you are running Docker on your machine, you have Minikube by default. 

**NOTE:** Minikube is only for testing simple Kubernetes applications in the development environment. You can never use Minikube while in a production environment.

To check the version of Minikube that came with Docker, run this `minikube` command:

`minikube version`

To start a Minikube Kubernetes Cluster, run this `minikube` command:

`minikube start --driver=docker`

The command will create and start a Minikube Kubernetes Cluster as shown below:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_6.png)

We have started a Minikube Kubernetes Cluster on our local machine. Before we deploy the Dockerized Fast API application, we need to set up the Kubernetes CLI (kubectl).

## Getting Started With the Kubernetes CLI

Kubernetes CLI (kubectl) is a powerful command line tool interface that allows DevOps engineers to communicate with a Kubernetes Cluster from the terminal. It has commands for managing and deploying Dockerized applications to the Kubernetes Cluster. There are different ways of installing the Kubernetes CLI. You will install this tool using the following command:

choco install kubernetes-cli

To check if the installation was successful, run this `kubectl` command:

`kubectl version --client`

When you want to deploy an application to Minikube Kubernetes Cluster using kubectl, you will need to create a Kubernetes YAML file. Let’s talk about this file.

## The Kubernetes YAML File

The Kubernetes YAML file describes the resources for the Kubernetes application and pods in the Kubernetes Cluster. This file also configures the [Kubernetes Service](https://kubernetes.io/docs/concepts/services-networking/service/) for the application.

Kubectl will use this file to deploy the Dockerized Fast API application to Minikube. In the ‘fast-api’ folder, create a new file named ‘kubernetes.yaml’. Open the file and add the following code:

apiVersion: apps/v1

kind: Deployment

metadata:

  name: fast-api-deployment

spec:

  replicas: 2

  selector:

​    matchLabels:

​      app: fast-api

  template:

​    metadata:

​      labels:

​        app: fast-api

​    spec:

​      containers:

​      \- name: fast-api

​        image: bravinwasike/fast-api

​        resources:

​          limits:

​            memory: "128Mi"

​            cpu: "500m"

​        ports:

​        \- containerPort: 8000

\---

apiVersion: v1

kind: Service

metadata:

  name: fast-api-service

spec:

  selector:

​    app: fast-api

  ports:

  \- port: 8000

​    targetPort: 8000

  type: LoadBalancer

The ‘kubernetes.yaml’ is divided into two. The first part is known as ‘Deployment’ while the second is known as ‘Service’.

##### 1. Deployment

It configures the application pods and the resources of your deployed application. It will create two pods. 

Pods are replicas or instances of the deployed application. It will use the ‘bravinwasike/fast-api’ image from Docker Hub to create the Docker Container.

##### 2. Service

It configures the Kubernetes Service for the application. It uses the `LoadBalancer` Kubernetes Service for distributing traffic to the two container pods. 

Minikube will assign an External IP address to the Kubernetes Service. It will enable us to access the deployed Fast API application.

Let’s apply Kubectl to this file:

kubectl apply -f kubernetes.yaml

The command will create a Kubernetes Deployment named ‘fast-api-deployment’. It will also create a Kubernetes Service named ‘fast-api-service’. Let’s run the following commands to get the Kubernetes Deployment and Service:

##### 1. Kubernetes Deployment

Run the following command:

kubectl get deployment

It gives the following output:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_7.png)

##### 2. Kubernetes Service

Run the following command:

kubectl get service

It gives the following output:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_8.png)

## Accessing the Deployed FastAPI Application

To access the deployed FastAPI application, Minikube will assign an External IP address to the ‘fast-api-service’. You will type the External IP address in your web browser and access the application. To get the External IP address, run this ‘minikube’ command:

minikube service fast-api-service

Minikube has assigned an External IP address to the `fast-api-service` as shown below:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_9.png)

Type the assigned IP address in your web browser to access your Dockerized FastAPI application:

![img](https://sweetcode.io/wp-content/uploads/2022/12/Bravin-How-to-Dockerize-Fast-API_10.png)

Our Dockerized FastAPI application is running in the Minikube Kubernetes Cluster. We have successfully dockerized and deployed a Fast API application to Kubernetes.

## Conclusion

In this tutorial, you have learned how to Dockerize and Deploy a Fast API application to Kubernetes. We started by creating and running a Fast API application locally. Then, we wrote a Dockerfile and used it to build a Docker Image for the FastAPI application.

We used the Docker Image to launch a Docker Container. Finally, we deployed the Dockerized Fast API application to the Minikube Kubernetes. We used the Kubernetes CLI and the Kubernetes YAML to deploy the application.