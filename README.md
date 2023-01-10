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
go run main.go --port=3000 --mod=release

# build
go build
```

## golang



