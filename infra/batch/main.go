package main

import (
	"log"

	"github.com/Bo0km4n/d2d/infra/batch/db"
	"github.com/Bo0km4n/d2d/infra/batch/pubsub"
	nats "github.com/nats-io/go-nats"
)

var queueGroup string

func init() {
	db.InitMySQL()
	pubsub.InitNATS("localhost", "4222", "d2d")
	queueGroup = "queue_1"
}

func main() {
	for {
		pubsub.PubsubConn.QueueSubscribe(queueGroup, func(m *nats.Msg) {
			log.Printf("message: %s", string(m.Data))
		})
	}
}
