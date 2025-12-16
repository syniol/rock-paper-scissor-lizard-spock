package game

const MinimumPlayerNameLength = 3
const MaximumPlayerNameLength = 12

type Player struct {
	Name   string
	Choice Choice
}

func NewPlayer() *Player {
	return &Player{}
}
