![build](https://github.com/freeshineit/gin_serve/workflows/build/badge.svg)


## Use

```bash
# 1. create docker network
docker network create my_net

# 初始化 mysql , redis 等 （前提是已经安装docker） 
# 2. initialization docker container, mysql, redis...
docker-compose -f db-docker-compose.yaml up -d
# or (Makefile)
make db

# 3. hot run development
# https://github.com/cosmtrek/air Live reload for Go apps
air

## or
## 3. run development
# http://localhost:8080
# http://localhost:8081
# http://localhost:8082
# http://localhost:8080/swagger/index.html
go run main.go

## Build Production
# Run make build to compile app.
make build

# help
./bin/app --help

# run production 
./bin/app --mode=release

# docker deploy app
docker build -t gin_serve:latest .
# http://localhost:3000
# http://localhost:3001
# http://localhost:3002
# run docker
docker run -it -p 3000:3000 -p 3001:3001 -p 3002:3002 --rm --net my_net --name gin_serve_api_service gin_serve:latest

# or (Makefile)
make serve

# generate api docs
swag init -g ./main.go

http://localhost:8080/swagger/index.html

#
go mod tidy

```

## Feature 

- [x] RESTful API 
- [x] JWT authentication
- [x] Email login/register
- [x] Swagger api document
- [ ] [Sentry](https://sentry.io/welcome/)



