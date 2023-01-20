FROM golang
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . /usr/app
ENV GOPROXY="https://goproxy.io"
RUN go build  -i -v -ldflags '-w -s' -o ./bin/app ./cmd/app.go
COPY ./public /usr/app/bin/public
COPY ./templates /usr/app/bin/templates
# COPY ./config/config.yaml /usr/app/bin/config
EXPOSE 3000
EXPOSE 3001
ENTRYPOINT ./bin/app --port=3000 --mode=release --proxy-port=3001 