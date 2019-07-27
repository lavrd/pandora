package broker

import (
	"crypto/tls"

	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/encoders/protobuf"

	"pandora/pkg/conf"
	"pandora/pkg/utils/errors"
)

const (
	SubMasterBlock = "MASTER_BLOCK"
	SubCertBlock   = "CERT_BLOCK"
	SubCert        = "CERT"
)

// Broker
type Broker struct {
	conn *nats.EncodedConn
}

// New connect to broker
func New(endpoint, user, password string) (*Broker, error) {
	cert, err := tls.LoadX509KeyPair(conf.Conf.TLS.Cert, conf.Conf.TLS.Key)
	if err != nil {
		return nil, errors.WithStack(err)
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
		return nil, errors.WithStack(err)
	}

	conn, err := nats.NewEncodedConn(c, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, errors.WithStack(err)
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
		return errors.WithStack(err)
	}
	return nil
}

// Publish bind send channel to subject
func (b *Broker) Publish(subject string, ch interface{}) error {
	if err := b.conn.BindSendChan(subject, ch); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
