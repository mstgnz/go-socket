package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	port := "3333"

	socket := NewSocket()
	http.Handle("/", websocket.Handler(socket.handler))

	done := make(chan bool)
	go func() {
		panic(http.ListenAndServe(":"+port, nil))
	}()
	log.Println("Started Websocket on :3333")
	<-done
}

/*
	use commands in browser console tab
	let socket = new WebSocket("ws://localhost:3333")
	socket.onmessage = (event) => {console.log("message: ", event.data)}
	socket.send("test message")
*/
