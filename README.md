# http_server_for_docker
Test HTTP server repository to learn Docker 


# Build a docker image
To align with yaml file later, lets build a image ithinkusc/hello-world-http-server
```
docker build . -t ithinkusc/hello-world-http-server
```

# Start on Minikube
After starting the minikube,

```
kubectl apply -f deployment.yaml
```

The pod can start but not successful because it cannot find the right container image to pull.
To fix this, we need to make the image accessible in the minikube

## Option 1: Load the image into minikube
```
minikube image load ithinkusc/hello-world-http-server
```

## Optioin 2: Directly build the image in minikube
```
minikube image build -t ithinkusc/hello-world-http-server
```

After the fix, re-apply the yaml file the server Pod is ready.

You can port forward the local port 8090 to the 8091 of the deployment using

```
kubectl port-forward deployment/hello-world-server 8090:8090
```

Learn more about loading local image to minikube from https://levelup.gitconnected.com/two-easy-ways-to-use-local-docker-images-in-minikube-cd4dcb1a5379 