package game

type Choice string

const (
	Rock    Choice = "Rock"
	Paper   Choice = "Paper"
	Scissor Choice = "Scissor"
	Lizard  Choice = "Lizard"
	Spock   Choice = "Spock"
)

var Rules = map[Choice]map[Choice]string{
	Scissor: {
		Paper:  "Cuts Paper",
		Lizard: "Decapitates Lizard",
	},
	Paper: {
		Rock:  "Covers Rock",
		Spock: "Disproves Spock",
	},
	Rock: {
		Scissor: "Crushes Scissor",
		Lizard:  "Crushes Lizard",
	},
	Lizard: {
		Paper: "Eats Paper",
		Spock: "Poisons Spock",
	},
	Spock: {
		Scissor: "Smashes Scissor",
		Rock:    "Vaporizes Rock",
	},
}
