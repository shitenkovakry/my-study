package main

import (
	"context"
	"log"
	"my-study/mygrpc"

	"google.golang.org/grpc"
)

func main() {
	// Устанавливаем соединение с gRPC сервером
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}
	defer conn.Close()

	//создаем grpc клиент
	roiuteMessageClient := mygrpc.NewRoiuteMessageClient(conn)

	response, err := roiuteMessageClient.SayHello(context.TODO(), &mygrpc.HelloRequest{
		Name:    "this is my name",
		Message: "this is my message",
	})

	if err != nil {
		panic(err)
	}

	log.Print(response.Message)
}
