package main

import (
	"fmt"
	"net/http"
)

const (
	address = ":8080"
)

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello, client")
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(address, nil)
}
