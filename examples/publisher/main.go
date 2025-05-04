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
