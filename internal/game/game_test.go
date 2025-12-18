package game

import "testing"

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

	// todo: could be a matrix of all choices in iteration for all possible outcomes
	sut.Players[0].Name = "John"
	sut.Players[0].Choice = Rock

	sut.Players[1].Name = "Alice"
	sut.Players[1].Choice = Paper

	result := sut.Start()

	if result.Winner.Name != sut.Players[1].Name {
		t.Errorf("Winner should be Alice but %s has won", sut.Start().Winner.Name)
	}

	if result.Reason != "Paper Covers Rock" {
		t.Errorf("Reason should be %s", sut.Start().Reason)
	}

	if sut.Scoreboard.storage[sut.Players[1].Name] != 1 {
		t.Errorf("Scoreboard should have a winning score record for Alice")
	}
}
