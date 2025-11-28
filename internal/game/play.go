package game

import (
	"fmt"

	"interview-rock-paper/internal/game/scoreboard"
)

// Play accepts two player as a parameter
// When there is no error and result is nil, it's a draw
func Play(playerOne, playerTwo *Player) *scoreboard.Score {
	if playerOne.Choice() == playerTwo.Choice() {
		return &scoreboard.Score{
			Status: scoreboard.PlayStatusDraw,
			Reason: fmt.Sprintf("both players chose %s", playerOne.Choice()),
		}
	}

	if reason, has := Rules[playerOne.Choice()][playerTwo.Choice()]; has {
		return &scoreboard.Score{
			Status: scoreboard.PlayStatusWin,
			Winner: playerOne.ID,
			Reason: fmt.Sprintf("%s %s", playerOne.Choice(), reason),
		}
	}

	if reason, has := Rules[playerTwo.Choice()][playerOne.Choice()]; has {
		return &scoreboard.Score{
			Status: scoreboard.PlayStatusWin,
			Winner: playerTwo.ID,
			Reason: fmt.Sprintf("%s %s", playerTwo.Choice(), reason),
		}
	}

	return &scoreboard.Score{
		Status: scoreboard.PlayStatusDraw,
		Reason: fmt.Sprintf("unexpected matching draw players chose %s %s", playerOne.Choice(), playerTwo.Choice()),
	}
}
