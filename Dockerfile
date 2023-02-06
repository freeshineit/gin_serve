FROM golang
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . /usr/app
ENV GOPROXY="https://goproxy.io"
ENV GIN_MODE=release
COPY ./Makefile /usr/app/
RUN make build
# COPY ./public /usr/app/bin/public
# COPY ./templates /usr/app/bin/templates
# COPY ./config/config-production.toml /usr/app/bin/config
EXPOSE 3000
EXPOSE 3001
EXPOSE 3002
ENTRYPOINT ./bin/app --mode=release --config ./config/config-production.toml