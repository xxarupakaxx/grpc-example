package server

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/xxarupakaxx/grpc-example/gen/pb"
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
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrusEnty)))

	pb.RegisterMatchingServiceServer(server, NewMatchService())

	reflection.Register(server)

	go func() {
		log.Println("start grpc server ")

		if err = server.Serve(lis); err != nil {
			log.Println("failed to start server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping grpc server")
	server.GracefulStop()
}
