package main

import (
	"fmt"
	"net/http"
)

const (
	address = ":8080"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello, client")
}

func main() {
	http.HandleFunc("/", helloHandler)

	http.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "this is information about the server")
	})

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}
