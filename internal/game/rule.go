package game

type Choice string

const (
	Rock     Choice = "Rock"
	Paper    Choice = "Paper"
	Scissors Choice = "Scissors"
	Lizard   Choice = "Lizard"
	Spock    Choice = "Spock"
)

var Rules map[Choice]map[Choice]string = map[Choice]map[Choice]string{
	Scissors: {
		Paper:  "Cuts Paper",
		Lizard: "Decapitates Lizard",
	},
	Paper: {
		Rock:  "Covers Rock",
		Spock: "Disproves Spock",
	},
	Rock: {
		Scissors: "Crushes Scissors",
		Lizard:   "Crushes Lizard",
	},
	Lizard: {
		Paper: "Eats Paper",
		Spock: "Poisons Spock",
	},
	Spock: {
		Scissors: "Smashes Scissors",
		Rock:     "Vaporizes Rock",
	},
}
