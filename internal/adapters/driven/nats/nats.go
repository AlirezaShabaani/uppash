package nats_handler

import (
	"context"
	"github.com/nats-io/nats.go"
)

type MessageBrokerHandler interface {
	Publisher(ctx context.Context, subject string, data []byte) error
	Subscriber(ctx context.Context, subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error)
}

type natsHandler struct {
	natsConn *nats.Conn
}

func NewMessageBrokerHandler(natsConn *nats.Conn) MessageBrokerHandler {
	return &natsHandler{natsConn: natsConn}
}

func (nh *natsHandler) Publisher(ctx context.Context, subject string, data []byte) error {
	err := nh.natsConn.Publish(subject, data)
	if err != nil {
		return err
	}
	return nil
}

func (nh *natsHandler) Subscriber(ctx context.Context, subject string, handler func(msg *nats.Msg)) (*nats.Subscription, error) {
	sub, err := nh.natsConn.Subscribe(subject, handler)
	if err != nil {
		return nil, err
	}
	return sub, nil
}
