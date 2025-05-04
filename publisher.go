package nexor

import (
	"context"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func (n *Nexor) Publish(ctx context.Context, subject string, msg proto.Message, opts ...nats.PubOpt) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = n.js.Publish(subject, data, opts...)
	return err
}
