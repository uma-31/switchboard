package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprint(writer, "Hello, I'm manager!")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	http.HandleFunc("/hello", hello)

	server.ListenAndServe()
}
