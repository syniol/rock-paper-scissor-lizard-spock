package game

import (
	"fmt"
	"math/rand"
)

type State string

const (
	StatePlay  State = "playing"
	StateReset State = "reset"
	StateExit  State = "exit"
)

type Opponent string

const (
	OpponentComputer Opponent = "computer"
	OpponentHuman    Opponent = "human"
)

type Game struct {
	Players    [2]*Player
	Scoreboard *Scoreboard
	State      State
	Opponent   Opponent
}

func NewGame() *Game {
	return &Game{
		Scoreboard: NewScoreboard(),
		Players: [2]*Player{
			NewPlayer(),
			NewPlayer(),
		},
		State:    StatePlay,
		Opponent: OpponentComputer,
	}
}

func (g *Game) Start() *Score {
	if g.Opponent == OpponentComputer {
		g.Players[1].Name = "Computer"

		computerSelection := rand.Intn(len(Rules))
		count := 0
		for weaponName, _ := range Rules {
			if count == computerSelection {
				g.Players[1].Choice = weaponName
			}
			count++
		}
	}

	if g.Players[0].Choice == g.Players[1].Choice {
		return &Score{
			Status: PlayStatusDraw,
			Reason: fmt.Sprintf("both players picked %s", g.Players[0].Choice),
		}
	}

	if reason, has := Rules[g.Players[0].Choice][g.Players[1].Choice]; has {
		g.Scoreboard.SetScore(g.Players[0].Name)

		return &Score{
			Status: PlayStatusWin,
			Winner: g.Players[0],
			Reason: fmt.Sprintf("%s %s", g.Players[0].Choice, reason),
		}
	}

	if reason, has := Rules[g.Players[1].Choice][g.Players[0].Choice]; has {
		g.Scoreboard.SetScore(g.Players[1].Name)
		return &Score{
			Status: PlayStatusWin,
			Winner: g.Players[1],
			Reason: fmt.Sprintf("%s %s", g.Players[1].Choice, reason),
		}
	}

	return &Score{
		Status: PlayStatusDraw,
		Reason: fmt.Sprintf("unexpected matching draw players picked %s %s", g.Players[0].Choice, g.Players[1].Choice),
	}
}
