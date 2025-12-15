package game

import (
	"github.com/google/uuid"
)

const MinimumPlayerNameLength = 3
const MaximumPlayerNameLength = 12

type Player struct {
	ID     PlayerID
	Name   string
	Choice Choice
}

func NewPlayer() *Player {
	return &Player{
		ID: PlayerID(uuid.NewString()),
	}
}
