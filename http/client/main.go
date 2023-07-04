package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

func main() {
	//создаем новый http клиент
	httpClient := &http.Client{}

	//создаем http запрос
	request, err := http.NewRequest("GET", "http://mystudy.com", nil)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "can not create request"))

		return
	}

	// отправляем запрос с помощью клиента
	response, err := httpClient.Do(request)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "can not send response"))

		return
	}
	defer response.Body.Close()

	// читаем тело ответа
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "can not read answer from the server"))

		return
	}

	// выводим тело ответа
	fmt.Println(string(body))
}
