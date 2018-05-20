package main

import (
	"net/http"
	"io"
	"log"
)

func main() {
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/pong", pongHandler)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("Error:", err)
	}
}
func pongHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Pong...")
}
func pingHandler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Ping...")
}
