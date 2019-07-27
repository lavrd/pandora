FROM golang:1.12.7-alpine3.9
RUN apk update && apk upgrade && \
    apk add --no-cache git
WORKDIR /pandora
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./cmd/kit/kit ./cmd/kit/kit.go
ENTRYPOINT ["./cmd/kit/kit", "-c", "/pandora/contrib/docker/conf.yml"]
