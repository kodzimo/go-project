package main

import (
	"final-project-kodzimo-hashing/internal/hashing"
	"final-project-kodzimo-hashing/internal/storage"
	pb "final-project-kodzimo-shared/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	redisClient, err := storage.ConnectToRedis()
	if err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	hashingService := hashing.NewHashingService(redisClient)

	/*
		Этот код (ниже) создает gRPC сервер и регистрирует ваш Hashing Service на этом сервере.
		Затем он начинает слушать входящие запросы на порту 50051.
	*/

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHashingServer(s, &hashing.Server{HashingService: hashingService})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
