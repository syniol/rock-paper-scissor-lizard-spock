package game

import "interview-rock-paper/internal/game/scoreboard"

type Player struct {
	ID     scoreboard.PlayerID
	choice *Choice
}

// Choice returning a value on purpose to encapsulate memory address (idempotency)
func (p *Player) Choice() Choice {
	return *p.choice
}

func (p *Player) SetChoice(choice *Choice) *Player {
	p.choice = choice

	return p
}
