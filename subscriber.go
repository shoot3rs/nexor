// Package nexor provides a NATS client implementation with support for protobuf message handling.
// It simplifies the process of publishing and subscribing to messages using Protocol Buffers
// for message serialization and deserialization.
package nexor

import (
	"context"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

// ProtoHandler is a function type that defines the signature for handling protobuf messages.
// It processes a decoded protobuf message along with its NATS message context and returns an error if processing fails.
type ProtoHandler func(ctx context.Context, msg proto.Message, m *nats.Msg) error

// Subscriber interface defines the contract for types that can handle protobuf message subscriptions.
// It requires implementations to provide both the protobuf message type and a handler function.
type Subscriber interface {
	// ProtoMessage returns a new instance of the protobuf message type that this subscriber handles
	ProtoMessage() proto.Message
	// Handler processes the received protobuf message and its associated NATS message
	Handler(ctx context.Context, msg proto.Message, m *nats.Msg) error
}

// Subscribe sets up a subscription to a NATS subject with protobuf message handling.
// It automatically decodes incoming messages using the provided protobuf message factory
// and processes them with the specified handler.
//
// Parameters:
//   - subject: The NATS subject to subscribe to
//   - durable: The durable name for the subscription (for JetStream persistence)
//   - factory: A function that creates new instances of the protobuf message type
//   - handler: A function that processes decoded protobuf messages
//   - opts: Optional subscription options that override the default settings
//
// Default behavior:
//   - Uses durable subscriptions for message persistence
//   - Requires manual message acknowledgment
//   - Sets a 30-second acknowledgment timeout
//
// Returns:
//   - error: Returns an error if the subscription setup fails
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
	sub, err := n.js.Subscribe(subject, func(m *nats.Msg) {
		// Create a new instance of the protobuf message
		msg := factory()
		if err := proto.Unmarshal(m.Data, msg); err != nil {
			if n.cfg.Debug {
				log.Printf("❌ nexor: failed to decode protobuf: %v", err)
			}

			_ = m.Nak() // NACK to let NATS know we couldn't process the message
			return
		}

		// Call the handler to process the message
		if err := handler(context.Background(), msg, m); err != nil {
			if n.cfg.Debug {
				log.Printf("❌ nexor: handler error: %v", err)
			}

			_ = m.Nak() // NACK if the handler fails
			return
		}
	}, opts...)

	if err != nil {
		if n.cfg.Debug {
			log.Printf("❌ nexor: failed to subscribe to subject: %s: %v", subject, err)
		}
		return err
	}

	if n.cfg.Debug {
		log.Printf("🚀 nexor: successfully subscribed to subject: %s", sub.Subject)
	}

	return err
}
