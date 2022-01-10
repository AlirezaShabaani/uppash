package nats_handler

import (
	"github.com/nats-io/nats.go"
	"log"
)

func InitNats(natsUrl string) (nc *nats.Conn) {
	var err error
	nc, err = nats.Connect(natsUrl)
	if err != nil {
		log.Fatalf("can't establish connection to nats server err : %v", err.Error())
	}
	return
}

