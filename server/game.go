package server

import (
	"fmt"
	"github.com/xxarupakaxx/grpc-example/game"
	"github.com/xxarupakaxx/grpc-example/game/utils"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"sync"
)

type GameService struct {
	sync.RWMutex
	tic    map[int32]*game.TicTacToe
	client map[int32][]pb.GameService_PlayServer
}

func NewGameService() *GameService {
	return &GameService{
		tic:    make(map[int32]*game.TicTacToe),
		client: make(map[int32][]pb.GameService_PlayServer),
	}
}

func (g *GameService) Play(stream pb.GameService_PlayServer) error {
	for true {
		g.Lock()
		defer g.Unlock()
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		switch req.GetAction().(type) {
		case *pb.PlayRequest_Start:

			ga := g.tic[req.GetRoomID()]
			if ga == nil {
				ga = game.NewTicTacToe(game.UNKNOWN)
				g.tic[req.GetRoomID()] = ga
				g.client[req.GetRoomID()] = make([]pb.GameService_PlayServer, 0, 2)

			}
			g.client[req.GetRoomID()] = append(g.client[req.GetRoomID()], stream)

			if len(g.client[req.GetRoomID()]) == 2 {
				for _, s := range g.client[req.GetRoomID()] {
					err = s.Send(&pb.PlayResponse{
						Event: &pb.PlayResponse_Ready{
							Ready: &pb.PlayResponse_ReadyEvent{},
						},
					})
					if err != nil {
						return err
					}
				}
				fmt.Println("game start")
			}
		case *pb.PlayRequest_Playing:
			ga := g.tic[req.GetRoomID()]

			winner := ga.Logic()
			for _, s := range g.client[req.GetRoomID()] {
				err = s.Send(&pb.PlayResponse{
					Event: &pb.PlayResponse_Play{
						Play: &pb.PlayResponse_PlayEvent{
							Player: req.GetPlayer(),
							Board:  utils.ConvertToPbBoard(g.tic[req.GetRoomID()].Board),
						},
					},
				})
				if err != nil {
					return err
				}
				if winner != game.UNKNOWN {
					err = stream.Send(&pb.PlayResponse{
						Event: &pb.PlayResponse_Finish{
							Finish: &pb.PlayResponse_FinishEvent{
								Result: utils.ConvertToPbResult(game.Winner(winner, utils.ConvertToGamePiece(req.GetPlayer().GetPiece()))),
								Board:  utils.ConvertToPbBoard(g.tic[req.GetRoomID()].Board),
							},
						},
					})
					if err != nil {
						return err
					}

					delete(g.client, req.GetRoomID())
				}
			}

		}
	}
	return nil
}
