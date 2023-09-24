package types

type Players []Player

func (p *Players) AddPlayer(player Player) Player {
	*p = append(*p, player)
	return player
}

func (p *Players) FindPlayer(name string) *Player {
	for _, player := range *p {
		if player.Name == name {
			return &player
		}
	}
	return nil
}

func (p *Players) DelPlayer(name string) {
	for i, player := range *p {
		if player.Name == name {
			*p = append((*p)[:i], (*p)[i+1:]...)
			return
		}
	}
}
