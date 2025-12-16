package game

type PlayStatus string

const (
	PlayStatusWin  PlayStatus = "Win"
	PlayStatusDraw PlayStatus = "Draw"
)

type Score struct {
	Status PlayStatus
	Winner *Player
	Reason string
}

type Scoreboard struct {
	storage map[string]int
}

func NewScoreboard() *Scoreboard {
	return &Scoreboard{
		storage: make(map[string]int),
	}
}

func (sb *Scoreboard) SetScore(key string) {
	sb.storage[key] = sb.storage[key] + 1
}

func (sb *Scoreboard) Scoreboard() map[string]int {
	return sb.storage
}
