package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/charmbracelet/huh"
	"interview-rock-paper/internal/game"
)

func main() {
	newGame := game.NewGame()

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your gaming nickname?").
				Value(&newGame.Players[0].Name).
				Validate(func(name string) error {
					return newGame.Players[0].ValidateName()
				}),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	for newGame.State != game.StateExit {
		err = huh.NewForm(
			huh.NewGroup(
				huh.
					NewSelect[game.Choice]().
					Title("Pick your weapon").
					Options(
						huh.NewOption("Rock 🪨", game.Rock),
						huh.NewOption("Paper 📄", game.Paper),
						huh.NewOption("Scissor ✂️", game.Scissor),
						huh.NewOption("Lizard 🦎", game.Lizard),
						huh.NewOption("Spock 🖖", game.Spock),
					).Value(&newGame.Players[0].Choice),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}

		score := newGame.Start()

		fmt.Printf("You picked: %s\n", newGame.Players[0].Choice)
		fmt.Printf("Computer picked: %s\n", newGame.Players[1].Choice)
		if score.Winner != nil {
			fmt.Printf("The winner is: %s\n\n", score.Winner.Name)
		}

		if score.Winner == nil {
			fmt.Printf("%s %s\n\n", score.Status, score.Reason)
		}

		if newGame.Scoreboard.HasScore() {
			tableHeaderColour := color.New(color.FgWhite, color.BgMagenta, color.Bold).SprintfFunc()
			tableColumnColour := color.New(color.FgYellow, color.Bold, color.BlinkSlow).SprintfFunc()
			scoreboardTable := table.New("Score", "Name")
			scoreboardTable.
				WithHeaderFormatter(tableHeaderColour).
				WithFirstColumnFormatter(tableColumnColour).
				WithPadding(6)

			for playerName, playerScore := range newGame.Scoreboard.Scoreboard() {
				scoreboardTable.AddRow(playerScore, playerName)
			}
			scoreboardTable.Print()
			fmt.Println()
		}

		err = huh.NewForm(
			huh.NewGroup(
				huh.
					NewSelect[game.State]().
					Options(
						huh.NewOption("Continue", game.StatePlay),
						huh.NewOption("Reset Scoreboard", game.StateReset),
						huh.NewOption("Exit", game.StateExit),
					).Value(&newGame.State),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}

		if newGame.State == "reset" {
			newGame.Scoreboard.Reset()
			fmt.Println("Scoreboard has been reset\n")
		}

		if newGame.State == "exit" {
			fmt.Println("\nThank you for playing.\n")
			os.Exit(0)
		}
	}
}
