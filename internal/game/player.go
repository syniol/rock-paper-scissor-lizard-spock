package game

import (
	"github.com/google/uuid"
)

type Player struct {
	ID     PlayerID
	Name   string
	choice *Choice
}

func NewPlayer(name string) *Player {
	return &Player{
		ID:     PlayerID(uuid.NewString()),
		Name:   name,
		choice: nil,
	}
}

// Choice returning a value on purpose to encapsulate memory address (idempotency)
func (p *Player) Choice() Choice {
	return *p.choice
}

func (p *Player) SetChoice(choice Choice) *Player {
	p.choice = &choice

	return p
}
