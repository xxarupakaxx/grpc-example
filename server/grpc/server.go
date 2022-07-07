package main

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
	"github.com/xxarupakaxx/grpc-example/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen : %w", err)
	}

	logusLogger := logrus.New()
	logrusEnty := logrus.NewEntry(logusLogger)
	grpc_logrus.ReplaceGrpcLogger(logrusEnty)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrusEnty)))

	pb.RegisterMatchingServiceServer(s, server.NewMatchService())

	reflection.Register(s)

	go func() {
		log.Println("start grpc server ")

		if err = s.Serve(lis); err != nil {
			log.Println("failed to start server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server")
	s.GracefulStop()
}
