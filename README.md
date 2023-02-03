# k8s_go_api
A small repo to demonstrate deploying a dockerised Golang rest api 

About:
* Restful api uses Golang and Fiber framework. Air pkg used for hot reload in dev env.
* Project is fully dockerised for later k8s deployment
* DB for local run a postgres instance using GORM for ORM functionality. K8s deploy will ofc have a DB set up 
  outside of the pod
* Added templating from the GoFiber framework to experiment with it.
To run:
`docker-compose up`

`docker compose run --service-ports web bash`

go mod init github.com/Kay-Wilkinson/k8s_go_api

go get github.com/gofiber/fiber/v2

go run cmd/main.go -b 0.0.0.0

Go to localhost:3000 to view Index page

Hot reload through go air pkg. Config for air pkg in .air.toml
