# Rock, Paper, Scissor, Lizard, Spock in Go <sup>(GoLang)</sup>
![workflow](https://github.com/syniol/rock-paper-scissor-lizard-spock/actions/workflows/pipeline.yml/badge.svg)

<div align="center">
    <p align="center" style="width: 50%; float: left;"><a href="https://syniol.com" target="blank"><img alt="Single Player Demo" width="50%" src="https://github.com/syniol/rock-paper-scissor-lizard-spock/blob/main/docs/single-player-demo-optimised.gif?raw=true" /></a>
    <sup><br />Single Player Demo</sup></p>

<p align="center" style="width: 50%; float: right;"><a href="https://syniol.com" target="blank"><img alt="Two Players Demo" width="50%" src="https://github.com/syniol/rock-paper-scissor-lizard-spock/blob/main/docs/two-players-demo-optimised.gif?raw=true" /></a>
    <sup><br />Two Players Demo</sup></p>
</div>

<a href='https://www.youtube.com/watch?v=pIpmITBocfM' rel="noopener noreferrer" target='_blank'>The game</a> invented by Sam Kass with his wife Karen Bryla is played between two players. Each player chooses one of the five options:
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


## Motivation
I recently got asked to come in for a second interview at this awesome robotics company in London. I was super pumped about what they do.

However, the pair programming exercise could be pretty tough for me. I struggle with stress, and they switched up the technical challenge 
at the last minute. The engineer assigned to the exercise seemed really nice, though. I think he was a bit let down, but he had no idea 
about my ADD, which I always keep under wraps. I even flunked my motorbike test nine times before finally passing, and I never mentioned 
my condition to the examiner.

This was one of their take-home exercises that I started but never got to demo, and I ended up failing the pairing exercise. I wrapped it 
up just to save face, hoping the engineers at that company might see that I'm not as bad as I seem, haha.


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


### Dependencies
 * [CLI Forms](https://github.com/charmbracelet/huh)
 * [CLI Table Colour Format](https://github.com/fatih/color)
 * [CLI Table Display](https://github.com/rodaine/table)


#### Credits
Author: [Hadi Tajallaei (Genius Persian Child)](mailto:hadi@syniol.com)

Copyright &copy; 2025 Syniol Limited. All rights reserved.
