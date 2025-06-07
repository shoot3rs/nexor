package nexor

import (
	"context"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

type Requester interface {
	Request(ctx context.Context, subject string, req proto.Message, factory func() proto.Message, timeout time.Duration) (proto.Message, error)
}

// Request sends a protobuf message as a request and waits for a protobuf reply.
// - subject: The NATS subject to send the request to
// - req: The protobuf message to send
// - factory: A function that returns a new instance of the expected reply message
// - timeout: How long to wait for a response
func (n *nexor) Request(ctx context.Context, subject string, req proto.Message, factory func() proto.Message, timeout time.Duration) (proto.Message, error) {
	data, err := proto.Marshal(req)
	if err != nil {
		if n.cfg.Debug {
			log.Printf("❌ nexor: failed to marshal request: %v", err)
		}
		return nil, err
	}

	msg, err := n.conn.RequestWithContext(ctx, subject, data)
	if err != nil {
		if n.cfg.Debug {
			log.Printf("❌ nexor: request error: %v", err)
		}
		return nil, err
	}

	reply := factory()
	if err := proto.Unmarshal(msg.Data, reply); err != nil {
		if n.cfg.Debug {
			log.Printf("❌ nexor: failed to unmarshal response: %v", err)
		}
		return nil, err
	}

	return reply, nil
}
