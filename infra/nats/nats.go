package nats

import (
	"fmt"
	"log"

	nats "github.com/nats-io/go-nats"
)

// NATS //
type NATS struct {
	Subj string
	Conn *nats.Conn
}

// NATSConn handler
var NATSConn *NATS

// New //
func New(host, port, subj string) error {
	var natsConn NATS
	address := fmt.Sprintf("nats://%s:%s", host, port)
	nc, err := nats.Connect(address)
	if err != nil {
		return err
	}
	natsConn.Conn = nc
	natsConn.Subj = subj
	NATSConn = &natsConn
	return nil
}

// Publish //
func (n *NATS) Publish(m []byte) error {
	return n.Conn.Publish(n.Subj, m)
}

// QueueSubscribe //
func (n *NATS) QueueSubscribe(queue string) {
	n.Conn.QueueSubscribe(n.Subj, queue, func(m *nats.Msg) {
		log.Println(m)
	})
}
