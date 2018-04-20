package broker

import (
	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/log"
)

const (
	SubjectBlock = "block"
)

// Broker
type Broker struct {
	*nats.EncodedConn
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

// Close close connection with broker server
func (b *Broker) Close() {
	b.EncodedConn.Close()
}

// Subscribe subscribe to broker messages by subject
func (b *Broker) Subscribe(subject string, ch interface{}) error {
	_, err := b.BindRecvChan(subject, ch)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Publish publish message to broker by subject
func (b *Broker) Publish(subject string, message interface{}) error {
	err := b.EncodedConn.Publish(subject, message)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
