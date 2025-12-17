package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"

	"github.com/syniol/rock-paper-scissor-lizard-spock/internal/game"
)

func main() {
	newGame := game.NewGame()

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Please enter your gaming nickname?").
				Value(&newGame.Players[0].Name).
				Validate(func(name string) error {
					return newGame.Players[0].ValidateName()
				}),
		),
		huh.NewGroup(
			huh.NewSelect[game.Opponent]().
				Title("Do you wish to play against a computer?").
				Options(
					huh.NewOption("Yes", game.OpponentComputer),
					huh.NewOption("No", game.OpponentHuman),
				).Value(&newGame.Opponent),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	if newGame.Opponent == game.OpponentHuman {
		err = huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Second player, please enter your gaming nickname?").
					Value(&newGame.Players[1].Name).
					Validate(func(name string) error {
						return newGame.Players[1].ValidateName()
					}),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}
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

		if newGame.Opponent == game.OpponentHuman {
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
						).Value(&newGame.Players[1].Choice),
				),
			).Run()
			if err != nil {
				log.Fatal(err)
			}
		}

		score := newGame.Start()

		fmt.Printf(
			"%s picked: %s\n",
			newGame.Players[0].Name,
			newGame.Players[0].Choice,
		)
		fmt.Printf(
			"%s picked: %s\n",
			newGame.Players[1].Name,
			newGame.Players[1].Choice,
		)
		if score.Winner != nil {
			fmt.Printf(
				"The winner is: %s. %s\n\n",
				score.Winner.Name,
				score.Reason,
			)
		}

		if score.Winner == nil {
			fmt.Printf("%s %s\n\n", score.Status, score.Reason)
		}

		if newGame.Scoreboard.HasScore() {
			newGame.Scoreboard.Print()
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

		if newGame.State == game.StateReset {
			newGame.Scoreboard.Reset()
			fmt.Print("Scoreboard has been reset\n\n")
		}

		if newGame.State == game.StateExit {
			fmt.Print("\nThank you for playing.\n\n")
			os.Exit(0)
		}
	}
}
