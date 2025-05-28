package main

import (
	"context"
	"github.com/shoot3rs/nexor"
	v1 "github.com/shoot3rs/nexor/gen/shooters/nexor/v1"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

var (
	client = nexor.New("nats://localhost:4222")
)

func init() {
	client.Connect()
}

func main() {
	ctx := context.Background()
	resp, err := client.Request(ctx, "example.say.hello", &v1.SayHelloRequest{Name: "Joey"}, func() proto.Message { return &v1.SayHelloResponse{} }, 3*time.Second)
	if err != nil {
		log.Println("🚨 response error:", err)
	}

	response := resp.(*v1.SayHelloResponse)

	log.Println("==== 📣 Response==== :", response.GetMessage())
}
