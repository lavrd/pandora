package broker

import (
	"fmt"

	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/log"
)

// Broker
type Broker struct {
	*nats.EncodedConn
}

// Connect connect to broker server
func Connect(url string, port int) (*Broker, error) {
	c, err := nats.Connect(fmt.Sprintf("%s:%d", url, port))
	if err != nil {
		log.Error(err)
		return nil, err
	}

	conn, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
	}
	return &Broker{conn}, err
}

// Close close connection with broker server
func (b *Broker) Close() {
	b.EncodedConn.Close()
}

// Subscribe subscribe to broker messages by channel
func (b *Broker) Subscribe(channel string, ch interface{}) error {
	_, err := b.BindRecvChan(channel, ch)
	if err != nil {
		log.Error(err)
	}
	return err
}

// Publish publish message to broker by channel
func (b *Broker) Publish(channel string, message interface{}) error {
	err := b.Publish(channel, message)
	if err != nil {
		log.Error(err)
	}
	return err
}
