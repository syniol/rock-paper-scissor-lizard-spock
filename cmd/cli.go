package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"unicode/utf8"

	"github.com/charmbracelet/huh"

	"interview-rock-paper/internal/game"
)

func main() {
	err := huh.NewForm(
		huh.NewGroup(
			huh.
				NewSelect[string]().
				Options(
					huh.NewOption("New Game", "new"),
					huh.NewOption("Continue", "continue"),
				),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	player := game.NewPlayer()

	err = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter your gaming nickname?").
				Value(&player.Name).
				Validate(func(name string) error {
					if utf8.RuneCountInString(name) < game.MinimumPlayerNameLength {
						return fmt.Errorf("name must contain at least %d characters", game.MinimumPlayerNameLength)
					}

					if utf8.RuneCountInString(name) > game.MaximumPlayerNameLength {
						return fmt.Errorf("name must contain at most %d characters", game.MaximumPlayerNameLength)
					}

					return nil
				}),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	var exitRequested bool = false
	for exitRequested == false {
		err = huh.NewForm(
			huh.NewGroup(
				huh.
					NewSelect[game.Choice]().
					Title("Pick your weapon").
					Options(
						huh.NewOption("Rock 🪨", game.Rock),
						huh.NewOption("Paper 📄", game.Paper),
						huh.NewOption("Scissor ✂️", game.Scissors),
						huh.NewOption("Lizard 🦎", game.Lizard),
						huh.NewOption("Spock 🖖", game.Spock),
					).Value(&player.Choice),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}

		computer := game.NewPlayer()
		computer.Name = "Computer"
		computerSelection := rand.Intn(len(game.Rules))
		count := 0
		for weaponName, _ := range game.Rules {
			if count == computerSelection {
				computer.Choice = weaponName
			}
			count++
		}

		score := game.Play(computer, player)

		fmt.Printf("Computer picked: %s\n", computer.Choice)
		if score.Winner != nil {
			fmt.Printf("The winner is: %s\n\n", score.Winner.Name)
		}

		if score.Winner == nil {
			fmt.Printf("%s %s\n\n", score.Status, score.Reason)
		}

		err = huh.NewForm(
			huh.NewGroup(
				huh.
					NewSelect[bool]().
					Title("").
					Options(
						huh.NewOption("Continue", false),
						huh.NewOption("Exit", true),
					).Value(&exitRequested),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	if exitRequested {
		fmt.Println("\nThank you for playing.")
		os.Exit(0)
	}
}
