# pandora

Blockchain-based decentralized platform for issuing, viewing, and verifying certificates

### Docker images
ArangoDB - `docker run -p 8529:8529 -e ARANGO_ROOT_PASSWORD=root arangodb`

NATS - `docker run -d -p 4222:4222 -p 8222:8222 -p 6222:6222 --name gnatsd nats:latest --user root --pass root`
