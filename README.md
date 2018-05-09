# pandora
Blockchain-based decentralized platform for issuing, viewing, and verifying certificates

### Docker images
##### ArangoDB
`docker run -d -p 8529:8529 -e ARANGO_ROOT_PASSWORD=root --name arangodb arangodb`
##### NATS
`docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 --name gnatsd nats:latest --user root --pass root`
### Proto
`protoc -I ./pkg/membership/pb --go_out=plugins=grpc:./pkg/membership/pb ./pkg/membership/pb/membership.proto`
`protoc -I ./pkg/discovery/pb --go_out=plugins=grpc:./pkg/discovery/pb ./pkg/discovery/pb/discovery.proto`
`protoc -I ./pkg/master/pb --go_out=plugins=grpc:./pkg/master/pb ./pkg/master/pb/master.proto`
