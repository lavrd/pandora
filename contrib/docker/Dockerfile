FROM golang:1.10.2

WORKDIR /go/src/pandora/cmd/kit
COPY . ../../

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

WORKDIR /go/src/pandora/cmd/kit
RUN go build

CMD ["./kit"]
