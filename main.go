package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	port := "3000"

	http.Handle("/", http.FileServer(http.Dir("public")))

	socket := NewSocket()
	http.Handle("/ws", websocket.Handler(socket.handler))

	done := make(chan bool)
	go func() {
		panic(http.ListenAndServe(":"+port, nil))
	}()
	log.Println("Started Websocket on :" + port)
	<-done
}
