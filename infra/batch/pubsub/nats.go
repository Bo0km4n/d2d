package pubsub

import (
	"fmt"

	nats "github.com/nats-io/go-nats"
)

// PubsubConn //
var PubsubConn *Pubsub

// Pubsub //
type Pubsub struct {
	Subj string
	Conn *nats.Conn
}

// InitNATS //
func InitNATS(host, port, subj string) error {
	var natsConn Pubsub
	address := fmt.Sprintf("nats://%s:%s", host, port)
	nc, err := nats.Connect(address)
	if err != nil {
		return err
	}
	natsConn.Conn = nc
	natsConn.Subj = subj
	PubsubConn = &natsConn
	return nil
}

// Publish //
func (n *Pubsub) Publish(m []byte) error {
	return n.Conn.Publish(n.Subj, m)
}

// QueueSubscribe //
func (n *Pubsub) QueueSubscribe(queue string, f func(m *nats.Msg)) {
	n.Conn.QueueSubscribe(n.Subj, queue, f)
}
