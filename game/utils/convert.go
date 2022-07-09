package utils

import (
	"github.com/xxarupakaxx/grpc-example/game"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
)

func ConvertToGamePlayer(pbPlayer *pb.Player) *game.Player {
	return &game.Player{
		PlayerID: int(pbPlayer.GetPlayerID()),
		Piece:    ConvertToGamePiece(pbPlayer.GetPiece()),
	}
}

func ConvertToGamePiece(pbPiece *pb.Piece) game.XO {
	switch pbPiece.Xo {
	case pb.Piece_X:
		return game.X
	case pb.Piece_O:
		return game.O
	default:
		return game.UNKNOWN

	}
}

func ConvertToPbPiece(piece game.XO) *pb.Piece {
	switch piece {
	case game.X:
		return &pb.Piece{Xo: pb.Piece_X}
	case game.O:
		return &pb.Piece{Xo: pb.Piece_O}
	default:
		return &pb.Piece{Xo: pb.Piece_UNKNOWN}
	}
}

func ConvertToPbPlayer(player *game.Player) *pb.Player {
	return &pb.Player{
		PlayerID: int32(player.PlayerID),
		Piece:    ConvertToPbPiece(player.Piece),
	}
}
