package client

import (
	"context"
	"fmt"
	"github.com/xxarupakaxx/grpc-example/game"
	"github.com/xxarupakaxx/grpc-example/game/utils"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"sync"
)

type Game struct {
	sync.RWMutex
	me   *game.Player
	room *game.Room
	game *game.Board
}

func (g *Game) MatchingGame(c pb.MatchingServiceClient) error {
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
			g.me = utils.ConvertToGamePlayer(res.GetPlayer())
			fmt.Println("マッチング相手が見つかりました")
			return nil
		}
	}
	return nil
}
