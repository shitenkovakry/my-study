package main

import (
	"log"
	"my-study/mygrpc"
	"my-study/roiute"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	address = ":8080"
)

func main() {
	// создание tcp прослушивателя
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Print(errors.Wrapf(err, "can not listen"))
	}

	//создание нового сервера grpc
	serverOfGrpc := grpc.NewServer()
	server := roiute.New()

	// регистрация службы на сервере
	mygrpc.RegisterRoiuteMessageServer(serverOfGrpc, server)

	log.Println("server started, listening on ", address)

	// запуск сервера grpc для прослушивания соединений
	if err := serverOfGrpc.Serve(listen); err != nil {
		log.Print(errors.Wrapf(err, "failed to serve"))
	}
}
