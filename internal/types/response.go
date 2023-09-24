package types

type Response struct {
	Type     string    `json:"type"`
	Message  string    `json:"message"`
	Player   Player    `json:"player"`
	Players  []Player  `json:"players"`
	Messages []Message `json:"messages"`
}
