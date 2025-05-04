package nexor

import (
	"github.com/nats-io/nats.go"
	"time"
)

type Nexor struct {
	conn *nats.Conn
	js   nats.JetStreamContext
}

// New initializes a new Nexor client with optional NATS options.
// Usage: nexor.New("nats://localhost:4222")
func New(url string, opts ...nats.Option) (*Nexor, error) {
	if len(opts) == 0 {
		opts = []nats.Option{
			nats.Name("NexorClient"),
			nats.MaxReconnects(5),
			nats.ReconnectWait(2 * time.Second),
		}
	}

	conn, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}

	js, err := conn.JetStream()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Nexor{conn: conn, js: js}, nil
}

// Close safely closes the NATS connection.
func (n *Nexor) Close() {
	if n.conn != nil && !n.conn.IsClosed() {
		n.conn.Close()
	}
}

// JetStream exposes the underlying JetStream context
// so that microservices can create/manage streams and consumers.
func (n *Nexor) JetStream() nats.JetStreamContext {
	return n.js
}
