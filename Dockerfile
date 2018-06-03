FROM golang:1.10.2

WORKDIR /go/src/github.com/spacelavr/pandora

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

WORKDIR /go/src/github.com/spacelavr/pandora/cmd/kit
RUN go build
RUN mv ./kit ../../

WORKDIR /go/src/github.com/spacelavr/pandora
ENTRYPOINT ["./kit", "-c", "./contrib/docker/conf.yml"]
