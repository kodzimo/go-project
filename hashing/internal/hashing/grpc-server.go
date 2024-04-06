package hashing

import (
	"context"

	pb "final-project-kodzimo-shared/proto"
)

type Server struct {
	pb.HashingServer
	HashingService *HashingService
}

func (s *Server) CheckHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	return s.HashingService.CheckHash(ctx, in)
}

func (s *Server) GetHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	return s.HashingService.GetHash(ctx, in)
}

func (s *Server) CreateHash(ctx context.Context, in *pb.HashRequest) (*pb.HashResponse, error) {
	return s.HashingService.CreateHash(ctx, in)
}

// Вынесено в main.go
//
// func main() {
// 	lis, err := net.Listen("tcp", ":50051")
// 	if err != nil {
// 		log.Fatalf("failed to listen: %v", err)
// 	}
// 	s := grpc.NewServer()
// 	pb.RegisterHashingServer(s, &Server{})
// 	if err := s.Serve(lis); err != nil {
// 		log.Fatalf("failed to serve: %v", err)
// 	}
// }

/*
Этот код создает gRPC сервер и регистрирует ваш Hashing Service на этом сервере.
Затем он начинает слушать входящие запросы на порту 50051.

Методы CheckHash, GetHash и CreateHash в grpc-server.go являются частью реализации gRPC сервера.
Они служат “обработчиками” для соответствующих RPC вызовов, которые могут быть сделаны клиентом
(в вашем случае, Gateway Service).

Когда Gateway Service делает вызов CheckHash, GetHash или CreateHash, gRPC сервер получает
этот вызов и перенаправляет его на соответствующий обработчик в grpc-server.go. Этот обработчик
затем выполняет необходимую логику и возвращает результат обратно через gRPC сервер к Gateway Service.

Однако, в вашем случае, вы правильно заметили, что эти методы также должны быть реализованы
в HashingService. В идеале, HashingService должен содержать основную бизнес-логику вашего приложения,
а gRPC сервер должен просто перенаправлять вызовы от Gateway Service к HashingService.

Таким образом, вместо того чтобы напрямую реализовывать логику в grpc-server.go, вы можете делегировать
эту работу HashingService.

Здесь HashingService - это поле в структуре server, и каждый обработчик просто вызывает соответствующий
метод в HashingService. Это позволяет вам разделить логику вашего приложения (в HashingService) и логику
вашего gRPC сервера (в grpc-server.go), что делает ваш код более чистым и легким для понимания.
*/
