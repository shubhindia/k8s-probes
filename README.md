# K8s probes
## This repository contains the sample code which I used to demo the k8s readiness and liveness probes.

### Pre-requisites
- A working k8s cluster

### Steps to deploy the application
- Clone the repository
- Run the following command to deploy the golang application
```
kubectl apply -f manifests/golang-deployment.yaml
```
- Run the following command to deploy the python application
```
kubectl apply -f manifests/python-deployment.yaml
```