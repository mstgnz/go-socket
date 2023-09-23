package types

type Request struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Player  Player `json:"player"`
}
