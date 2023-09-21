package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

type Socket struct {
	name  string
	conns map[*websocket.Conn]bool
}

func NewSocket() *Socket {
	return &Socket{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Socket) handler(ws *websocket.Conn) {

	defer func(ws *websocket.Conn) {
		s.disconnect(ws)
	}(ws)

	s.name = ws.Request().URL.Query().Get("name")
	if len(s.name) == 0 {
		s.name = fmt.Sprintf("Guest%v", len(s.conns))
	}
	msg := fmt.Sprintf("%v connected", s.name)

	log.Println(msg)
	s.broadcast([]byte(msg))

	s.conns[ws] = true
	s.read(ws)
}

func (s *Socket) read(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println("Read error: ", err.Error())
			continue
		}
		s.broadcast(buf[:n])
	}
}

func (s *Socket) broadcast(b []byte) {
	for ws := range s.conns {
		if s.conns[ws] {
			go func(ws *websocket.Conn) {
				if _, err := ws.Write(b); err != nil {
					log.Println("Broadcast Error: ", err.Error())
				}
			}(ws)
		}
	}
}

func (s *Socket) disconnect(ws *websocket.Conn) {
	msg := fmt.Sprintf("%s disconnected", s.name)
	log.Println(msg)
	s.conns[ws] = false
	s.broadcast([]byte(msg))
	delete(s.conns, ws)
	_ = ws.Close()
}
