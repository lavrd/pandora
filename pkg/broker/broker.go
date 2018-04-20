package broker

import (
	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/log"
)

const (
	// SubjectBlock subject for send and read message about block
	SubjectBlock = "block"
	// SubjectCertificate subject for send and read message about certificate
	SubjectCertificate = "certificate"
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

// Subscribe bind receive channel to subject
func (b *Broker) Subscribe(subject string, ch interface{}) error {
	_, err := b.BindRecvChan(subject, ch)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Publish bind send channel to subject
func (b *Broker) Publish(subject string, ch interface{}) error {
	err := b.BindSendChan(subject, ch)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
