FROM golang:1.19
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . /usr/app
ENV GOPROXY="https://goproxy.io"
ENV GIN_MODE=release
RUN make build
EXPOSE 3000
EXPOSE 3001
EXPOSE 3002
ENTRYPOINT ./bin/app --mode=release --config ./config/config-production.toml