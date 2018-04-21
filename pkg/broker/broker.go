package broker

import (
	"time"

	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/log"
)

const (
	SubjectNewBlock       = "new_block"
	SubjectNewCertificate = "new_certificate"
	SubjectLastBlock      = "last_block"
	SubjectBlockchain     = "blockchain"
)

// Broker
type Broker struct {
	conn *nats.EncodedConn
}

// Connect connect to broker server
func Connect(endpoint string) (*Broker, error) {
	c, err := nats.Connect(endpoint)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	conn, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Broker{conn}, nil
}

// Reply reply to subject
func (b *Broker) Reply(subject string, f func() (interface{}, error)) error {
	if _, err := b.conn.Subscribe(subject, func(subject, reply string, msg interface{}) {
		if body, err := f(); err == nil {
			if err := b.conn.Publish(reply, body); err != nil {
				log.Error(err)
			}
		} else {
			log.Error(err)
		}
	}); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Request request by subject
func (b *Broker) Request(subject string, message, data interface{}) error {
	if err := b.conn.Request(subject, message, data, 1*time.Second); err != nil {
		log.Error(err)
		return err
	}
	return nil
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
