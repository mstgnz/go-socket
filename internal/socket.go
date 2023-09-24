// Package internal /*
/*
## License
This project is licensed under the MIT Licence. Refer to https://github.com/mstgnz/go-socket/blob/main/LICENSE for more information.
*/
package internal

import (
	"fmt"
	"io"

	"github.com/mstgnz/go-socket/types"
	"golang.org/x/net/websocket"
)

var players types.Players
var request types.Request
var messages types.Messages

type Socket struct {
	connections map[*websocket.Conn]bool
}

func NewSocket() *Socket {
	return &Socket{
		connections: make(map[*websocket.Conn]bool),
	}
}

func (s *Socket) Handler(ws *websocket.Conn) {

	defer func(ws *websocket.Conn) {
		s.disconnect(ws)
	}(ws)

	for {
		if err := websocket.JSON.Receive(ws, &request); err != nil {
			if err == io.EOF {
				break
			}
			HandleLog("read error", err)
			s.emit(ws, types.Response{Type: "error", Message: err.Error()})
			continue
		}
		switch request.Type {
		case "new":
			s.newHandle(ws)
		case "message":
			s.messageHandle(ws)
		case "animate":
			s.animateHandle(ws)
		case "name":
			s.nameHandle(ws)
		default:
			HandleLog("invalid type: "+request.Type, nil)
			s.emit(ws, types.Response{Type: "error", Message: "invalid type: " + request.Type})
		}
	}
}

func (s *Socket) emit(ws *websocket.Conn, response types.Response) {
	err := websocket.JSON.Send(ws, response)
	if err != nil {
		HandleLog("emit error", err)
	}
}

func (s *Socket) broadcast(response types.Response) {
	for ws := range s.connections {
		if s.connections[ws] {
			s.emit(ws, response)
			// TODO to many connection handle
			/*go func(ws *websocket.Conn) {
				s.emit(ws, response)
			}(ws)*/
		}
	}
}

func (s *Socket) disconnect(ws *websocket.Conn) {
	msg := fmt.Sprintf("%s disconnected", request.Player.Name)
	HandleLog(msg, nil)
	s.connections[ws] = false
	players.DelPlayer(request.Player.Name)
	delete(s.connections, ws)
	s.broadcast(types.Response{
		Type:    "disconnect",
		Message: msg,
		Player:  request.Player,
	})
	_ = ws.Close()
}

func (s *Socket) newHandle(ws *websocket.Conn) {
	player := players.AddPlayer(types.Player{
		Color:    RandomColor(),
		Name:     request.Player.Name,
		Position: types.Position{X: 0, Y: 0},
	})
	HandleLog(request.Player.Name+" connected", nil)
	s.emit(ws, types.Response{Type: "init", Message: "login successfully", Player: player, Players: players, Messages: messages})
	s.broadcast(types.Response{Type: request.Type, Message: "new player connected", Player: player})
}

func (s *Socket) messageHandle(ws *websocket.Conn) {
	messages = append(messages, types.Message{Name: request.Player.Name, Message: request.Message})
	if player := players.FindPlayer(request.Player.Name); player != nil {
		s.broadcast(types.Response{Type: request.Type, Message: request.Message, Player: *player})
	}
}

func (s *Socket) animateHandle(ws *websocket.Conn) {
	if player := players.FindPlayer(request.Player.Name); player != nil {
		player.Position = request.Player.Position
		s.broadcast(types.Response{Type: request.Type, Message: "animate", Player: *player})
	}
}

func (s *Socket) nameHandle(ws *websocket.Conn) {
	if player := players.FindPlayer(request.Player.Name); player != nil {
		player.Name = request.Player.Name
		s.broadcast(types.Response{Type: request.Type, Message: "change", Player: *player})
	}
}
