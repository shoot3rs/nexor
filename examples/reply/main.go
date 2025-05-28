package main

import (
	"context"
	"fmt"
	"github.com/shoot3rs/nexor"
	v1 "github.com/shoot3rs/nexor/gen/shooters/nexor/v1"
	"google.golang.org/protobuf/proto"
)

var (
	client = nexor.New("nats://localhost:4222")
)

func init() {
	client.Connect()
}

func main() {
	_ = client.Reply("example.say.hello",
		func() proto.Message { return &v1.SayHelloRequest{} },
		func(ctx context.Context, req proto.Message) (proto.Message, error) {
			request := req.(*v1.SayHelloRequest)
			return &v1.SayHelloResponse{Message: fmt.Sprintf("Hello %s", request.GetName())}, nil
		},
	)

	select {}
}
