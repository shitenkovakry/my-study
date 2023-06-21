package main

import (
	"log"
	"net"
	"time"

	"github.com/pkg/errors"
)

const (
	address  = "localhost:9091"
	stopWord = "stop"
)

func main() {
	// Подключение к серверу по адресу "localhost:8080"
	connection, err := net.Dial("tcp", address)
	if err != nil {
		panic(errors.Wrapf(err, "can not connection to server"))
	}

	defer connection.Close()

	messages := []string{
		"durik",
		"ondrys",
		// "miu",
		// "",
	}

	// Чтение данных от сервера
	for index := 0; index < len(messages); index++ {
		message := []byte(messages[index])
		wroteBytes, err := connection.Write(message)
		if err != nil {
			log.Print(errors.Wrapf(err, "can not write data"))
			return
		}

		log.Print("wrote ", wroteBytes, " bytes")
		time.Sleep(time.Second * 10)
	}
}
