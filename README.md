# Rock, Paper, Scissors, Lizard, Spock in Go (GoLang)

```json
{
  "status": "In Progress",
  "Estimated Finish Date": "17 December 2025"
}
```

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
```shell
make run
```


### Main Dependencies
 * [CLI Forms](https://github.com/charmbracelet/huh)
 * [UUID Player ID Assignment](https://github.com/google/uuid)


#### Credits
Author: [Hadi Tajallaei (Genius Persian Child)](mailto:hadi@syniol.com)

Copyright &copy; 2025 Syniol Limited. All rights reserved.
