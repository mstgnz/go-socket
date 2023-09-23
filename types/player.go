package types

type Player struct {
	Name     string   `json:"name"`
	Color    string   `json:"color"`
	Position Position `json:"position"`
}

type Position struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

func (u *Player) SetName(name string) {
	u.Name = name
}

func (u *Player) SetColor(color string) {
	u.Color = color
}

func (u *Player) SetPosition(position Position) {
	u.Position = position
}
