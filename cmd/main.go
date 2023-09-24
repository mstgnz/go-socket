package main

import (
	"log"
	"net/http"

	"github.com/mstgnz/go-socket/internal"
	"golang.org/x/net/websocket"
)

func main() {
	port := "3000"

	http.Handle("/", http.FileServer(http.Dir("public")))

	socket := internal.NewSocket()
	http.Handle("/ws", websocket.Handler(socket.Handler))

	done := make(chan bool)
	go func() {
		panic(http.ListenAndServe(":"+port, nil))
	}()
	log.Println("Started Websocket on :" + port)
	<-done
}
