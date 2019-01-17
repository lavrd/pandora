FROM golang:1.11.4-alpine3.8

WORKDIR /pandora

COPY . .

ENTRYPOINT ["./cmd/kit/kit", "-c", "../../contrib/docker/conf.yml"]
