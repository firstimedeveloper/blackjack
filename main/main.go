package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/firstimedeveloper/blackjack"
)

func main() {
	numOfPlayers, err := strconv.Atoi(blackjack.GetInput("Welcome to blackjack!\nEnter the number of players: "))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
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
