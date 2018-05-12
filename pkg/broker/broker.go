package broker

import (
	"github.com/nats-io/go-nats"
	"github.com/nats-io/go-nats/encoders/protobuf"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

const (
	SMasterBlock = "SMasterBLock"
	SCertBlock   = "SCertBlock"
)

// Broker
type Broker struct {
	conn *nats.EncodedConn
}

// Opts
type Opts struct {
	Endpoint string
	User     string
	Password string
}

// Connect connect to broker
func Connect(opts *Opts) (*Broker, error) {
	c, err := nats.Connect(
		opts.Endpoint,
		nats.UserInfo(opts.User, opts.Password),
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

// QSubscribe bind receive queue channel to subject
func (b *Broker) QSubscribe(subject, queue string, ch interface{}) error {
	if _, err := b.conn.BindRecvQueueChan(subject, queue, ch); err != nil {
		log.Error(err)
		return err
	}
	return nil
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
