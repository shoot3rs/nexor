# Nexor
A shooter package for handling events with nats jetstream

Example usage for publishing an event

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/shoot3rs/nexor"
	v1 "github.com/shoot3rs/nexor/gen/shooters/nexor/v1"
	"log"
	"strings"
	"time"
)

func main() {
	// Initialize the event bus
	bus, err := nexor.New("nats://localhost:4222")
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	js := bus.JetStream()

	// Create a stream (e.g. product.*)
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "PRODUCT_EVENTS",
		Subjects: []string{"product.*"},
		Storage:  nats.FileStorage,
	})
	if err != nil && !strings.Contains(err.Error(), "stream name already in use") {
		log.Fatalf("❌ Failed to create stream: %v", err)
	}

	for {
		// Create a ProductCreated event
		event := &v1.ProductCreated{
			Id:         uuid.NewString(),
			Name:       faker.Word(),
			SupplierId: uuid.NewString(),
			CreatedAt:  faker.UnixTime(),
		}

		// Publish the event
		if err := bus.Publish(context.Background(), "product.created", event); err != nil {
			log.Fatalf("Failed to publish event: %v", err)
		}

		fmt.Println("🚀 Event published successfully! 🚀")

		time.Sleep(time.Duration(10) * time.Second)
	}
}
```

Example usage for subscribing to an event

```go
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

```