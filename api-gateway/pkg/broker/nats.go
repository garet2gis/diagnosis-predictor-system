package broker

import (
	"github.com/nats-io/nats.go"
)

type NatsBroker struct {
	connection *nats.Conn
}

func newNatsBroker(connection *nats.Conn) *NatsBroker {
	return &NatsBroker{connection: connection}
}

func (n *NatsBroker) Subscribe(subj string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return n.connection.Subscribe(subj, cb)
}

func (n *NatsBroker) Publish(subj string, data []byte) error {
	return n.connection.Publish(subj, data)
}

func (n *NatsBroker) Drain() error {
	return n.connection.Drain()
}

func (n *NatsBroker) Close() {
	n.connection.Close()
}
