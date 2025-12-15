package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/charmbracelet/huh"

	"interview-rock-paper/internal/game"
)

func main() {
	var playerChoice game.Choice

	err := huh.NewForm(
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
				).Value(&playerChoice),
		),
	).Run()
	if err != nil {
		log.Fatal(err)
	}

	computer := game.NewPlayer("Computer")
	computerSelection := rand.Intn(len(game.Rules))
	count := 0
	for weaponName, _ := range game.Rules {
		if count == computerSelection {
			computer.SetChoice(weaponName)
		}
		count++
	}

	player := game.NewPlayer("You")
	player.SetChoice(playerChoice)

	score := game.Play(computer, player)

	fmt.Printf("Computer picked: %s\n", computer.Choice())
	if score.Winner != nil {
		fmt.Printf("The winner is: %s\n", score.Winner.Name)
	}

	if score.Winner == nil {
		fmt.Printf("%s %s\n", score.Status, score.Reason)
	}
}
