![build](https://github.com/freeshineit/gin_serve/workflows/build/badge.svg)


## Use

```bash

# 初始化 mysql , redis 等 （前提是已经安装docker） 
# initialization docker container, mysql, redis...
docker-compose -f docker-compose.yaml up -d

# run development
# https://github.com/cosmtrek/air Live reload for Go apps
air

## development
# http://localhost:8080
# http://localhost:8081
# http://localhost:8082
# http://localhost:8080/swagger/index.html
go run cmd/main.go

# build
make build

# help
./bin/app --help

# run production 
./bin/app --mode=release

# docker deploy app
docker build -t gin_serve:v0.1 .
# run docker
docker run -it -p 3000:3000 -p 3001:3001 -p 3002:3002 --rm --name gin_serve gin_serve:v0.1


# generate api docs
swag init -g ./cmd/main.go

http://localhost:8080/swagger/index.html

```

## golang



