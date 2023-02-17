# k8s_go_api
A small repo to demonstrate deploying a dockerised Golang rest api 

About:
* Restful api uses Golang and Fiber framework. Air pkg used for hot reload in dev env.
* Project is fully dockerised for later k8s deployment
* DB for local run a postgres instance using GORM for ORM functionality. K8s deploy will ofc have a DB set up 
  outside of the pod
* Added templating from the GoFiber framework to experiment with it.
* Hot reload through go air pkg. Config for air pkg in .air.toml
* K8s pods hosted via Linode
**** To Run
  Via Docker:
`docker-compose up`

`docker compose run --service-ports web bash`

  Via Go cli:
```
go mod init github.com/Kay-Wilkinson/k8s_go_api

go get github.com/gofiber/fiber/v2

go run cmd/main.go -b 0.0.0.0

Go to localhost:3000 to view Index page
```

K8s:

`export KUBECONFIG=kubeconfig.yaml`
``
```
$ kubectl get nodes
>>
NAME                           STATUS   ROLES    AGE     VERSION
lke93517-141511-63efc4ab6f31   Ready    <none>   5m54s   v1.25.4
lke93517-141511-63efc4abcf6e   Ready    <none>   6m48s   v1.25.4
lke93517-141511-63efc4ac3041   Ready    <none>   5m1s    v1.25.4
```
```
$ kubectl cluster-info
Kubernetes control plane is running at https://cc04eddd-6458-4fd8-a1d8-0771c09a5363.eu-central-2.linodelke.net:443
KubeDNS is running at https://cc04eddd-6458-4fd8-a1d8-0771c09a5363.eu-central-2.linodelke.net:443/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
```

```
$ kubectl run {pod_name} --image={ref to docker image} --port=3000

$ kubectl run goapipod --image=k8s_go_api_web --port=3000

>> pod/goapipod created
```

```
$ kubectl get pods
```