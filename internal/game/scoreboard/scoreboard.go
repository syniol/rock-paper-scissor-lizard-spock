package scoreboard

import (
	"encoding/json"
	"errors"
	"os"
	"sync"

	"interview-rock-paper/pkg"
)

type PlayStatus string

const (
	PlayStatusWin  PlayStatus = "Win"
	PlayStatusDraw PlayStatus = "Draw"
)

type PlayerID string

type Score struct {
	Status PlayStatus
	Winner PlayerID
	Reason string
}

type Scoreboard struct {
	History        map[PlayerID]int `json:"history"`
	LastPlayStatus PlayStatus       `json:"lastPlayStatus"`
	LastWinner     PlayerID         `json:"lastWinner"`

	entity pkg.Entity
}

var scoreboardStorage *Scoreboard
var onceScoreboard sync.Once

func find(ID string) (sb *Scoreboard, err error) {
	onceScoreboard.Do(func() {
		contents, errReadFile := os.ReadFile(".game_history.json")
		if errReadFile != nil {
			// could use ! but for better code readability I used == false
			if errors.Is(errReadFile, os.ErrNotExist) == false {
				if errReadFile != nil {
					err = errReadFile
					return
				}
			}

			// todo: create a file initial setup
		}

		var scoreboardStored Scoreboard
		err = json.Unmarshal(contents, &scoreboardStored)
		if err != nil {
			return
		}

		sb = scoreboardStorage
		return
	})

	return scoreboardStorage, nil
}

func Persis(sb *Scoreboard) error {
	return nil
}

func Delete(sb *Scoreboard) error {
	// todo: modify scoreboardStorage
	// todo: save
	return nil
}
