package game

import (
	"fmt"
	"unicode/utf8"
)

const MinimumPlayerNameLength = 3
const MaximumPlayerNameLength = 12

type Player struct {
	Name   string
	Choice Choice
}

func NewPlayer(name string) *Player {
	return &Player{
		Name: name,
	}
}

func (player *Player) ValidateName() error {
	if utf8.RuneCountInString(player.Name) < MinimumPlayerNameLength {
		return fmt.Errorf(
			"name must contain at least %d characters",
			MinimumPlayerNameLength,
		)
	}

	if utf8.RuneCountInString(player.Name) > MaximumPlayerNameLength {
		return fmt.Errorf(
			"name must contain at most %d characters",
			MaximumPlayerNameLength,
		)
	}

	return nil
}
