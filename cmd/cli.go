package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"unicode/utf8"

	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/charmbracelet/huh"
	"interview-rock-paper/internal/game"
)

func main() {
	var scoreboard = game.NewScoreboard()
	player := game.NewPlayer()

	err := huh.NewForm(
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

	var inGameCommand string = "continue"
	for inGameCommand != "exit" {
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
			scoreboard.SetScore(score.Winner.Name)

			fmt.Printf("The winner is: %s\n\n", score.Winner.Name)
		}

		if score.Winner == nil {
			fmt.Printf("%s %s\n\n", score.Status, score.Reason)
		}

		if scoreboard.HasScore() {
			tableHeaderColour := color.New(color.FgWhite, color.BgMagenta, color.Bold).SprintfFunc()
			tableColumnColour := color.New(color.FgYellow, color.Bold, color.BlinkSlow).SprintfFunc()
			scoreboardTable := table.New("Score", "Name")
			scoreboardTable.
				WithHeaderFormatter(tableHeaderColour).
				WithFirstColumnFormatter(tableColumnColour).
				WithPadding(6)

			for playerName, playerScore := range scoreboard.Scoreboard() {
				scoreboardTable.AddRow(playerScore, playerName)
			}
			scoreboardTable.Print()
			fmt.Println()
		}

		err = huh.NewForm(
			huh.NewGroup(
				huh.
					NewSelect[string]().
					Options(
						huh.NewOption("Continue", "continue"),
						huh.NewOption("Reset Scoreboard", "reset"),
						huh.NewOption("Exit", "exit"),
					).Value(&inGameCommand),
			),
		).Run()
		if err != nil {
			log.Fatal(err)
		}

		if inGameCommand == "reset" {
			scoreboard.Reset()
			fmt.Println("Scoreboard has been reset")
		}

		if inGameCommand == "exit" {
			fmt.Println("\nThank you for playing.")
			os.Exit(0)
		}
	}
}
