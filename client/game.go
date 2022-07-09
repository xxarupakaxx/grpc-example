package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/xxarupakaxx/grpc-example/game/utils"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	finish = false
	start  = false
)

func (g *Game) Play(client pb.GameServiceClient) error {
	stream, err := client.Play(context.Background())
	if err != nil {
		return err
	}

	defer stream.CloseSend()

	go func() {
		if err = g.send(stream); err != nil {
			log.Println(err)
			return
		}
	}()

	if err = g.receive(stream); err != nil {
		return err
	}

	return nil
}

func (g *Game) send(stream pb.GameService_PlayClient) error {
	for true {
		g.RLock()
		if finish {
			g.RUnlock()
			return nil
		} else if start {
			err := stream.Send(&pb.PlayRequest{
				RoomID: int32(g.room.RoomID),
				Player: utils.ConvertToPbPlayer(g.me),
				Action: &pb.PlayRequest_Start{
					Start: &pb.PlayRequest_StartAction{},
				},
			})
			if err != nil {
				log.Println("failed to send ", err)
				return err
			}
			g.RUnlock()

			for true {
				g.RLock()
				if start {
					g.RUnlock()
					fmt.Printf("対戦見つかったね")
					break
				}
				g.RUnlock()
				fmt.Println("対戦相手が見つかるまで待とうね")
				time.Sleep(time.Second * 1)
			}
		} else {
			g.RUnlock()
			stdin := bufio.NewScanner(os.Stdin)
			stdin.Scan()

			fmt.Println("どこに置く？")
			number := ParseText(stdin.Text())
			g.game.Line[number] = g.me.Piece

			g.Lock()
			g.game.DisplayBoard(g.me.Piece)
			g.Unlock()

			err := stream.Send(&pb.PlayRequest{
				RoomID: int32(g.room.RoomID),
				Player: utils.ConvertToPbPlayer(g.me),
				Action: &pb.PlayRequest_Playing{
					Playing: &pb.PlayRequest_PlayAction{
						PlaceNumber: number,
					},
				},
			})
			if err != nil {
				log.Println(err)
				return err
			}
			ch := make(chan struct{})
			go func(ch chan struct{}) {
				fmt.Println("")
				for i := 0; i < 5; i++ {
					fmt.Printf("%d秒間止まります \n", 5-i)
					time.Sleep(1 * time.Second)
				}
				fmt.Println("")
				ch <- struct{}{}
			}(ch)
			<-ch
		}
	}
	return nil
}

func (g *Game) receive(stream pb.GameService_PlayClient) error {
	for true {
		res, err := stream.Recv()
		if err != nil {
			return err
		}

		g.Lock()
		switch res.GetEvent().(type) {
		case *pb.PlayResponse_Ready:
			start = true
			g.game.DisplayBoard(g.me.Piece)
		case *pb.PlayResponse_Play:
			piece := utils.ConvertToGamePiece(res.GetPlay().GetPlayer().GetPiece())
			if piece != g.me.Piece {
				g.game.DisplayBoard(piece)
				fmt.Print("どこに置きますか　0~8")
			}
		case *pb.PlayResponse_Finish:
			finish = true
			fmt.Println("")
		default:

		}

		g.Unlock()
		return nil
	}

	g.Unlock()

	return nil
}
func ParseText(t string) int32 {
	t = strings.TrimSpace(t)
	number, _ := strconv.Atoi(t)
	return int32(number)
}
