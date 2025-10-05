# Quick start

When Clone this project after open the terminal or command line on your computer and cd to path of project and used command
```
docker compose up

or

using container run is a background on os
docker compose up -d --build
```

# How to generate swagger ui

You have to ssh to docker container first
```
is a using ssh connect to container
docker exec -it <container name or container id> sh
```

When connect to container success and used command to install swagger cmi
```
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$PATH:$(go env GOPATH)/bin
```

After use command below to generated swagger ui
```
swag init -g main.go --output packages/doc
```



All Endpoint of services
```
service -> http://127.0.0.1:3004
swagger ui -> http://localhost:3004/swagger/index.html
```
