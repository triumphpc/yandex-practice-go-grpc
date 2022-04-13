package main

import (
	"fmt"
	"github.com/triumphpc/yandex-practice-go-grpc/pkg/adder"
	pb "github.com/triumphpc/yandex-practice-go-grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// определяем порт для сервера
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}
	// создаём gRPC-сервер без зарегистрированной службы
	s := grpc.NewServer()
	// регистрируем сервис
	pb.RegisterUsersServer(s, &adder.UsersServer{})

	fmt.Println("сервер gRPC начал работу")
	// получаем запрос gRpc
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
