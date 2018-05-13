# pandora
Blockchain-based decentralized platform for issuing, viewing, and verifying certificates

### Docker images
##### ArangoDB
`docker run -d -p 8529:8529 -e ARANGO_ROOT_PASSWORD=root --name arangodb arangodb`
##### NATS
`docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats -v ./contrib:/tls nats:latest --user root --pass root --tls --tlscert /tls/cert.pem --tlskey /tls/key.pem`
### Proto
`protoc -I ./pkg/pb --go_out=plugins=grpc:./pkg/pb ./pkg/pb/pb.proto`
