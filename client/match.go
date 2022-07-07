package client

import (
	"context"
	"fmt"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
)

func MatchingGame(c pb.MatchingServiceClient) error {
	stream, err := c.MatchStreams(context.Background())
	if err != nil {
		return err
	}

	defer stream.CloseSend()

	for true {
		res, err := stream.Recv()
		if err != nil {
			return err
		}

		if res.GetStatus() == pb.MatchResponse_WAITING {
			fmt.Println("マッチング相手を探している")
		}
		if res.GetStatus() == pb.MatchResponse_MATCH {
			fmt.Println("マッチング相手が見つかりました")
			return nil
		}
	}
	return nil
}
