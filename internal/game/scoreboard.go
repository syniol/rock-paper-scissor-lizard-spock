package game

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

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

func (sb *Scoreboard) HasScore() bool {
	return len(sb.storage) > 0
}

func (sb *Scoreboard) Reset() {
	sb.storage = make(map[string]int)
}

func (sb *Scoreboard) Print() {
	tableHeaderColour := color.New(color.FgWhite, color.BgMagenta, color.Bold).SprintfFunc()
	tableColumnColour := color.New(color.FgYellow, color.Bold, color.BlinkSlow).SprintfFunc()
	scoreboardTable := table.New("Score", "Name")
	scoreboardTable.
		WithHeaderFormatter(tableHeaderColour).
		WithFirstColumnFormatter(tableColumnColour).
		WithPadding(6)

	for playerName, playerScore := range sb.Scoreboard() {
		scoreboardTable.AddRow(playerScore, playerName)
	}
	scoreboardTable.Print()
	fmt.Println()
}
