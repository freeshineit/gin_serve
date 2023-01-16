![build](https://github.com/freeshineit/go_python_serve/workflows/build/badge.svg)


## Use

```bash

# 初始化 mysql , redis 等 （前提是docker已经安装并启动） 前提是已经安装docker
docker-compose -f docker-compose.yaml up -d

# run development
# https://github.com/cosmtrek/air Live reload for Go apps
air

## development
# http://localhost:3000
go run cmd/app.go --port=3000

# build
make build

# help
./bin/app --help

# run production 
./bin/app --port=3000 --mode=release

# docker deploy app
docker build -t go_app:v0.1 .
# run docker
docker run -it -p 3000:3000 --rm --name go_app go_app:v0.1
```

## golang



