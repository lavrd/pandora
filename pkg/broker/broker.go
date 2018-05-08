package broker

import (
	"time"

	"github.com/nats-io/go-nats"
	"github.com/spacelavr/pandora/pkg/utils/log"
)

const (
	SCAccount = "SCAccount"
	SFAccount = "SFAccount"
	SCBlock   = "SCBlock"
	SNBlock   = "SNBlock"

	QCBlock = "QCBlock"
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

	conn, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &Broker{conn}, nil
}

// Reply reply to subject
func (b *Broker) Reply(subject string, handler func(m *nats.Msg)) error {
	if _, err := b.conn.Subscribe(subject, handler); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (b *Broker) SendReply(reply string, data interface{}) error {
	if err := b.conn.Publish(reply, data); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// Request request by subject
func (b *Broker) Request(subject string, message, data interface{}) error {
	if err := b.conn.Request(subject, message, data, time.Second*1); err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Close close connection with broker server
func (b *Broker) Close() {
	b.conn.Close()
}

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
