package broker

import (
	"crypto/tls"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/encoders/protobuf"

	"pandora/pkg/conf"
	"pandora/pkg/utils/log"
)

const (
	SUB_MASTER_BLOCK = "SUB_MASTER_BLOCK"
	SUB_CERT_BLOCK   = "SUB_CERT_BLOCK"
	SUB_CERT         = "SUB_CERT"
)

// Broker
type Broker struct {
	conn *nats.EncodedConn
}

// New connect to broker
func New(endpoint, user, password string) (*Broker, error) {
	cert, err := tls.LoadX509KeyPair(conf.Conf.TLS.Cert, conf.Conf.TLS.Key)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	tlsConfig := &tls.Config{
		ServerName:         endpoint,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{cert},
	}

	c, err := nats.Connect(
		endpoint,
		nats.UserInfo(user, password),
		nats.Secure(tlsConfig),
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	conn, err := nats.NewEncodedConn(c, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Broker{conn}, nil
}

// Close close connection with broker server
func (b *Broker) Close() {
	b.conn.Close()
}

// Subscribe bind receive channel to subject
func (b *Broker) Subscribe(subject string, ch interface{}) error {
	if _, err := b.conn.BindRecvChan(subject, ch); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Publish bind send channel to subject
func (b *Broker) Publish(subject string, ch interface{}) error {
	if err := b.conn.BindSendChan(subject, ch); err != nil {
		log.Error(err)
		return err
	}
	return nil
}
