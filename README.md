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
# http://localhost:3000
go run cmd/app.go --port=3000 --proxy-port=3001 

# build
make build

# help
./bin/app --help

# run production 
./bin/app --port=3000 --mode=release --proxy-port=3001 

# docker deploy app
docker build -t gin_serve:v0.1 .
# run docker
docker run -it -p 3000:3000 --rm --name gin_serve gin_serve:v0.1


# generate api docs
swag init -g ./cmd/app.go

http://localhost:8080/swagger/index.html

```

## golang



