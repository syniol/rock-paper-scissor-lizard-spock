package game

import (
	"fmt"
	"testing"
)

func TestNewGame(t *testing.T) {
	sut := NewGame()

	if sut.Opponent != OpponentComputer {
		t.Errorf("Initial Opponent should be %s unless is changed", OpponentComputer)
	}

	if sut.State != StatePlay {
		t.Errorf("Initial State should be %s unless is changed", StatePlay)
	}

	if sut.Scoreboard.HasScore() == true {
		t.Errorf("Initial scoreboard shoul be empty")
	}
}

func TestTwoPlayers(t *testing.T) {
	sut := NewGame()
	sut.Opponent = OpponentHuman

	if sut.Opponent != OpponentHuman {
		t.Errorf("Opponent should be %s", OpponentHuman)
	}

	for choice, opposeChoice := range Rules {
		sut.Players[0].Name = "John"
		sut.Players[0].Choice = choice

		for opponentChoice, reason := range opposeChoice {
			sut.Players[1].Name = "Alice"
			sut.Players[1].Choice = opponentChoice

			result := sut.Start()

			if result.Winner.Name != sut.Players[0].Name {
				t.Errorf("Winner should be Alice but %s has won", result.Winner.Name)
			}

			if result.Reason != fmt.Sprintf("%s %s", sut.Players[0].Choice, reason) {
				t.Errorf("Reason should include %s", reason)
			}

			if sut.Scoreboard.storage[sut.Players[0].Name] != 1 {
				t.Errorf("Scoreboard should have a winning score record for Alice")
			}

			sut.Scoreboard.Reset()
		}
	}
}

func TestSinglePlayer(t *testing.T) {
	sut := NewGame()

	sut.Players[0].Name = "John"
	sut.Players[0].Choice = Rock

	score := sut.Start()

	if len(score.Status) == 0 {
		t.Error("Expected a Status")
	}

	if len(score.Reason) == 0 {
		t.Error("Expected a reason")
	}
}
