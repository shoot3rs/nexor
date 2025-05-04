package nexor

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

type ProtoHandler func(ctx context.Context, msg proto.Message, m *nats.Msg) error

// Subscribe subscribes to a subject, listens for messages, and invokes the handler
// subject: The subject to listen to
// durable: Durable name for this subscription (ensures message delivery even if service restarts)
// factory: A function to create a new instance of the expected protobuf message
// handler: A function that processes the message once received
// opts: Additional options for the subscription (e.g., queue groups)
func (n *Nexor) Subscribe(
	subject string,
	durable string,
	factory func() proto.Message,
	handler ProtoHandler,
	opts ...nats.SubOpt,
) error {
	// Default options for the subscription
	defaultOpts := []nats.SubOpt{
		nats.Durable(durable),          // Use a durable subscription
		nats.ManualAck(),               // Acknowledge messages manually
		nats.AckWait(30 * time.Second), // Wait for ack within 30 seconds
	}

	// Merge any user-defined options with the default options
	opts = append(defaultOpts, opts...)

	// Subscribe to the subject with the provided options
	_, err := n.js.Subscribe(subject, func(m *nats.Msg) {
		// Create a new instance of the protobuf message
		msg := factory()
		if err := proto.Unmarshal(m.Data, msg); err != nil {
			log.Printf("❌ nexor: failed to decode protobuf: %v", err)
			_ = m.Nak() // NACK to let NATS know we couldn't process the message
			return
		}

		// Call the handler to process the message
		if err := handler(context.Background(), msg, m); err != nil {
			log.Printf("❌ nexor: handler error: %v", err)
			_ = m.Nak() // NACK if the handler fails
			return
		}
	}, opts...)

	return err
}
