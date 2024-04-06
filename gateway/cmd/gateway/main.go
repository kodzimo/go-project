package main

import (
	"final-project-kodzimo-gateway/internal/gateway"
	pb "final-project-kodzimo-shared/proto"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	/*
		В этом коде мы создаем новый экземпляр GatewayService, регистрируем обработчики HTTP для каждого
		из методов, а затем запускаем HTTP-сервер, который слушает на порту 8080. Обратите внимание, что мы
		запускаем gRPC сервер в отдельной горутине, чтобы основной поток мог продолжить и запустить HTTP-сервер.
	*/

	// Создаем соединение с gRPC сервером
	conn, err := grpc.Dial("hashing-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Создаем новый Gateway Service
	gw := &gateway.GatewayService{
		HashingClient: pb.NewHashingClient(conn),
	}

	// Регистрируем обработчики HTTP
	http.HandleFunc("/checkhash", gw.CheckHashHandler)
	http.HandleFunc("/gethash", gw.GetHashHandler)
	http.HandleFunc("/createhash", gw.CreateHashHandler)

	// Запускаем HTTP-сервер
	log.Fatal(http.ListenAndServe(":8080", nil))
}
