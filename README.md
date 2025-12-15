# Rock, Paper, Scissors, Lizard, Spock in Go (GoLang)

[The game](https://www.youtube.com/watch?v=pIpmITBocfM) is played between two players. Each player chooses one of the five options:
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


__Status:__ In Progress

__Estimated Finish Date:__ 17th December 2025


```shell
make run
```


### Main Dependencies
 * [CLI Forms](https://github.com/charmbracelet/huh)
 * [UUID Player ID Assignment](https://github.com/google/uuid)


##### Credits
Author: [Hadi Tajallaei](mailto:hadi@syniol.com)
Copyright &copy; 2025 Syniol Limited. All rights reserved.
