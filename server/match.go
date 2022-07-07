package server

import (
	"github.com/xxarupakaxx/grpc-example/game"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"sync"
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
	defer m.Unlock()
	count++
	player := &game.Player{
		PlayerID: count,
	}

	for _, room := range m.Room {
		if room.Player1 != nil {
			player.Piece = game.X
			room.Player2 = player

			err := stream.Send(&pb.MatchResponse{
				RoomID:   int32(room.RoomID),
				Status:   pb.MatchResponse_WAITING,
				PlayerID: int32(count),
			})

			if err != nil {
				return err
			}

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

	err := stream.Send(&pb.MatchResponse{
		RoomID:   int32(maxRoomID),
		Status:   pb.MatchResponse_WAITING,
		PlayerID: int32(player.PlayerID),
	})
	if err != nil {
		return err
	}

	go func() {
		for true {
			m.RLock()
			player2 := r.Player2
			m.Unlock()
			if player2 != nil {
				err = stream.Send(&pb.MatchResponse{
					RoomID:   int32(r.RoomID),
					Status:   pb.MatchResponse_MATCH,
					PlayerID: int32(player.PlayerID),
				})
				if err != nil {
					return
				}
			}
		}
	}()

	return nil
}
