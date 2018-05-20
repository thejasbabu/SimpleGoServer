package main

import (
	"net/http"
	"io"
	"fmt"
)

func main() {
	addr := ":8090"
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)
	fmt.Printf("Starting server at %s", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}

func pongHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Pong...")
}

func pingHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Ping...")
}
