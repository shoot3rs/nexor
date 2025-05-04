package main

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/shoot3rs/nexor"
	v1 "github.com/shoot3rs/nexor/gen/shooters/nexor/v1"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	// Initialize the event bus
	bus, err := nexor.New("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	defer bus.Close()

	// Subscribe to the "product.created" event
	err = bus.Subscribe("product.created", "product_created_consumer", func() proto.Message {
		return &v1.ProductCreated{} // Factory method to create a specific event type
	}, func(ctx context.Context, msg proto.Message, m *nats.Msg) error {
		log.Println("🔥 event received via subject:", m.Subject)

		// Type asserts the message to a specific event type
		productCreatedEvent, ok := msg.(*v1.ProductCreated)
		if !ok {
			log.Printf("👻 Received an unknown message type")
			return errors.New("unknown message type")
		}

		// Handle the event
		log.Printf("🔥 Product Created: %v", productCreatedEvent.String())

		if err := m.Ack(); err != nil {
			log.Println("❌ Failed to acknowledge message:", err)
		}

		return nil
	})

	if err != nil {
		log.Fatalf("❌ Failed to subscribe to product.created event: %v", err)
	}

	// Keep the main function running to receive events
	log.Println("🚀 waiting for events...")
	select {}
}
