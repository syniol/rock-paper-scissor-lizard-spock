# Rock, Paper, Scissor, Lizard, Spock in Go <sup>(GoLang)</sup>

<div align="center">
    <p align="center" style="width: 50%; float: left;"><a href="https://syniol.com" target="blank"><img alt="Single Player Demo" width="100%" src="https://github.com/syniol/rock-paper-scissor-lizard-spock/blob/main/docs/single-player-demo-optimised.gif?raw=true" /></a>
    <sup><br />Single Player Demo</sup></p>

<p align="center" style="width: 50%; float: right;"><a href="https://syniol.com" target="blank"><img alt="Two Players Demo" width="100%" src="https://github.com/syniol/rock-paper-scissor-lizard-spock/blob/main/docs/two-players-demo-optimised.gif?raw=true" /></a>
    <sup><br />Two Players Demo</sup></p>
</div>

<a href='https://www.youtube.com/watch?v=pIpmITBocfM' rel="noopener noreferrer" target='_blank'>The game</a> is played between two players. Each player chooses one of the five options:
- **Rock**
- **Paper**
- **Scissors**
- **Lizard**
- **Spock**

The winner is determined by the following rules:

| **Choice**   | **Wins Against** | **Reason**                       |
|--------------|------------------|----------------------------------|
| **Scissors** | Paper, Lizard    | Cuts Paper, Decapitates Lizard   |
| **Paper**    | Rock, Spock      | Covers Rock, Disproves Spock     |
| **Rock**     | Scissors, Lizard | Crushes Scissors, Crushes Lizard |
| **Lizard**   | Paper, Spock     | Eats Paper, Poisons Spock        |
| **Spock**    | Scissors, Rock   | Smashes Scissors, Vaporizes Rock |

If both players choose the same option, the game results in a **tie**.


## Features
- **Interactive Gameplay**: Players can select their choice, and the winner is determined based on the rules.
- **Clear Visual Feedback**: Winning and losing outcomes are displayed in an engaging and intuitive way.
- **Scoreboard**: Tracks the points of the user and the computer across multiple rounds.
- **Data Persistence**: Retains the game state and scoreboard.
- **Restart**: Allows the user to restart the game, clearing the scoreboard and resetting the game state.


## Up and Running
You need to have a Go installation for version `1.24`.

For Unix based operating systems you could use:
```shell
make run
```

For Windows, please use the following command:
```shell
go run cmd\cli.go
```


### Main Dependencies
 * [CLI Forms](https://github.com/charmbracelet/huh)
 * [CLI Table Colour Format](https://github.com/fatih/color)
 * [CLI Table Display](https://github.com/rodaine/table)


#### Credits
Author: [Hadi Tajallaei (Genius Persian Child)](mailto:hadi@syniol.com)

Copyright &copy; 2025 Syniol Limited. All rights reserved.
