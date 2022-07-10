package server

import (
	"github.com/xxarupakaxx/grpc-example/game"
	"github.com/xxarupakaxx/grpc-example/game/utils"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"sync"
	"time"
)

type MatchService struct {
	Room map[int]*game.Room
	sync.RWMutex
}

func NewMatchService() *MatchService {
	return &MatchService{Room: make(map[int]*game.Room)}
}

var count = 0

func (m *MatchService) MatchStreams(stream pb.MatchingService_MatchStreamsServer) error {
	m.Lock()

	count++
	player := &game.Player{
		PlayerID: count,
	}

	for _, room := range m.Room {
		if room.Player1 != nil {
			player.Piece = game.X
			room.Player2 = player

			err := stream.Send(&pb.MatchResponse{
				RoomID: int32(room.RoomID),
				Status: pb.MatchResponse_WAITING,
				Player: utils.ConvertToPbPlayer(player),
			})
			if err != nil {
				return err
			}
			m.Unlock()

			return nil
		}
	}

	player.Piece = game.O
	maxRoomID := len(m.Room) + 1
	r := &game.Room{
		RoomID:  maxRoomID,
		Player1: player,
		Player2: nil,
	}

	m.Room[r.RoomID] = r
	m.Unlock()

	err := stream.Send(&pb.MatchResponse{
		RoomID: int32(maxRoomID),
		Status: pb.MatchResponse_WAITING,
		Player: utils.ConvertToPbPlayer(player),
	})
	if err != nil {
		return err
	}

	ch := make(chan struct{})
	go func(ch chan<- struct{}) {
		for true {
			m.Lock()
			player2 := r.Player2
			m.Unlock()
			if player2 != nil {
				err = stream.Send(&pb.MatchResponse{
					RoomID: int32(r.RoomID),
					Status: pb.MatchResponse_MATCH,
					Player: utils.ConvertToPbPlayer(player),
				})
				if err != nil {
					return
				}
				ch <- struct{}{}
				break
			}
			time.Sleep(5 * time.Second)
		}
	}(ch)

	select {
	case <-ch:

	}

	return nil
}
