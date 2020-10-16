# nus-hacker-demo

A very simple demo for application deployment and upgrade on Kubernetes.

## Requirement

- Install Docker Desktop (https://www.docker.com/products/docker-desktop) and enable Kubernetes cluster.
- Install kubectl command line tool.

## Step 1 build Docker images for application

For application v1

```
docker build -t nus-hacker-demo:v1 ./app-v1
```

For application v2

```
docker build -t nus-hacker-demo:v2 ./app-v2
```

## Step 2 deploy application v1

Create a namespace `nus-hacker`

```
cd kubernetes-manifest
kubectl apply -f namespace.yaml
```

Deploy Service `nus-hacker-demo`

```
kubectl apply -f service.yaml
```

Deploy v1 

```
kubectl apply -f deployment.yaml
```

Go to `localhost:30000` to visit the application

## Step 3 scale up the Deployment

Change the `replica` field from 1 to 2

```
spec:
  replicas: 2
```

Deploy again

```
kubectl apply -f deployment.yaml
```

## Step 4 release application v2

Change the `image` field to v2

```
      containers:
      - name: app
        image: nus-hacker-demo:v2
```

Deploy again

```
kubectl apply -f deployment.yaml
```
