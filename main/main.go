package main

import (
	"fmt"
	"strconv"

	"github.com/firstimedeveloper/blackjack"
)

func main() {
	fmt.Println("Welcome to Black Jack!")
	numOfPlayers := 0
	invalidInput := true
	for invalidInput {
		var err error
		numOfPlayers, err = strconv.Atoi(blackjack.GetInput("Enter the number of players: "))
		if err != nil {
			fmt.Println("Not a valid Input")
		} else {
			invalidInput = false
		}
	}

	gameOver := false
	var input string
	for !gameOver {
		blackjack.StartGame(numOfPlayers)
		input = blackjack.GetInput("Keep playing(y/n)? ")
		if input == "n" {
			gameOver = true
			fmt.Println("Goodbye!")
		}
	}
}
