package game

import (
	"fmt"
)

// Play accepts two player as a parameter
// When there is no error and result is nil, it's a draw
func Play(playerOne, playerTwo *Player) *Score {
	if playerOne.Choice == playerTwo.Choice {
		return &Score{
			Status: PlayStatusDraw,
			Reason: fmt.Sprintf("both players picked %s", playerOne.Choice),
		}
	}

	if reason, has := Rules[playerOne.Choice][playerTwo.Choice]; has {
		return &Score{
			Status: PlayStatusWin,
			Winner: playerOne,
			Reason: fmt.Sprintf("%s %s", playerOne.Choice, reason),
		}
	}

	if reason, has := Rules[playerTwo.Choice][playerOne.Choice]; has {
		return &Score{
			Status: PlayStatusWin,
			Winner: playerTwo,
			Reason: fmt.Sprintf("%s %s", playerTwo.Choice, reason),
		}
	}

	return &Score{
		Status: PlayStatusDraw,
		Reason: fmt.Sprintf("unexpected matching draw players picked %s %s", playerOne.Choice, playerTwo.Choice),
	}
}
