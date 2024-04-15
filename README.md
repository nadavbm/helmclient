### helmclient

### test in minikube

follow these steps in the same terminal

build the image locally:
```
docker build -t helmut .
docker image ls
```

run minikube:
```
eval $(minikube docker-env)
minikube start --vm-driver=virtualbox --disk-size=10g
```

set in `kubernetes/deploy.yaml` for the relevant image and set `imagePullPolicy: Never`

### build and push image
