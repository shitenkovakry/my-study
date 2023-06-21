package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pkg/errors"
)

const (
	address            = "localhost:9091"
	stopWordFromClient = "stop"
)

func handleConnection(connection net.Conn) {
	defer connection.Close()

	addressOfClient := connection.RemoteAddr().String()
	log.Print("connected with client:", addressOfClient)

	bufferWindow := make([]byte, 1024)

	for {
		readLen, err := connection.Read(bufferWindow)
		if err != nil {
			log.Print(errors.Wrapf(err, "can not read message from client"))

			return
		}

		// Преобразуем прочитанные данные в строку.
		data := string(bufferWindow[:readLen])
		fmt.Printf("received data from the client %s: %s\n", addressOfClient, data)

		if data == stopWordFromClient {
			log.Println("recieved ", stopWordFromClient, " from client. close the connection")
			connection.Close()

			return
		}

		// Отправляем ответ клиенту.
		response := "Привет, клиент!"
		_, err = connection.Write([]byte(response))
		if err != nil {
			log.Print(errors.Wrapf(err, "can not send data to client"))

			return
		}

		fmt.Printf("sent a response to the client %s: %s\n", addressOfClient, response)
	}
}

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(errors.Wrapf(err, "can not listen connection"))
	}
	defer listener.Close()

	fmt.Println("server listening at:", address)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(errors.Wrapf(err, "can not accept connection"))

			continue
		}

		go handleConnection(connection)
	}
}
