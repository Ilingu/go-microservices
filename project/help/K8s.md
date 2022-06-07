# minikube

`minikube start` || `minikube start --nodes=2` --> Start/Init K8s
`minikube stop` --> Stop K8s
`minikube status` --> Current status of services

`minikube dashboard` -> Launch K8s dashboard interface of running service

# kubectl

`kubectl get pods` -> Get Pods
`kubectl get svc` -> Get Service
`kubectl get deployments` -> Get deployed services
`kubectl delete deployments <apps_name>` -> Delete deployed services

`kubectl apply -f <k8s_folder>` --> Create service from ConfigFile
`kubectl logs <service_name>` -> See logs of service
