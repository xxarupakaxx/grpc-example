package client

import (
	"context"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
)

func Play(client pb.GameServiceClient) error {
	stream, err := client.Play(context.Background())
	if err != nil {
		return err
	}

	for true {
		stream.Send()
	}
}
