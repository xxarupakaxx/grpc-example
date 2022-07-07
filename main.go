package main

import (
	"github.com/xxarupakaxx/grpc-example/client"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to grpc server :%w", err)
	}
	defer conn.Close()
	err = client.MatchingGame(pb.NewMatchingServiceClient(conn))
	if err != nil {
		return
	}
}