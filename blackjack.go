package blackjack

import (
	"fmt"

	"github.com/firstimedeveloper/deck"
)

type Player []deck.Card

//StartGame starts the game of BlackJack
func StartGame(numOfPlayers int) {
	gameDeck := deck.New(deck.Shuffle)
	for _, card := range gameDeck {
		fmt.Println(card)
	}
	getInput("Press enter to play blackjack")

	players := make([]Player, numOfPlayers)
	var dealtCard deck.Card
	for round := 0; round < 2; round++ {
		for i := 0; i < numOfPlayers; i++ {
			//dealtCard is the last Card in the deck
			//which is the card that is being dealt
			dealtCard = gameDeck[len(gameDeck)-1]
			gameDeck = gameDeck[:len(gameDeck)-1]
			//append dealtcard to hand
			players[i] = append(players[i], dealtCard)
			//players[0][0]
			//players[1][0]

		}
	}
	//printing dealt cards, while not showing last card
	//seems unnecessarily complicated
	//TODO make this shorter or something.
	printPlayerHand(players)

	//testing purposes
	//printing the scores
	for i, player := range players {
		fmt.Printf("Player %d score: %d\n", i+1, getValueHand(player))
	}

	var isValid bool //false
	var choice string
	for isValid {
		choice = getInput("hit or stand?")
		if choice == "hit" || choice == "stand" {
			isValid = true
		}
	}
	switch choice {
	case "hit":
		//somefunc
	case "stand":
		//somefunc
	}

}

func printPlayerHand(players []Player) {
	for i, player := range players {
		for j, c := range player {
			if i != len(players)-1 || j != 1 {
				if i+1 != len(players) {
					fmt.Printf("Player %d card %d: %v\n", i+1, j+1, c)

				} else {
					fmt.Printf("Dealer's card %d: %v\n", j+1, c)
				}
			}
		}
	}
}

func getValueHand(d []deck.Card) int {
	var score int
	for _, c := range d {
		score += getValueCard(c)
	}

	return score
}

func getValueCard(c deck.Card) int {
	switch c.Rank {
	case deck.Jack, deck.Queen, deck.King:
		return 10
	case deck.Ace:
		return 11
	default:
		return int(c.Rank)
	}
}

func getInput(phrase string) string {
	fmt.Println(phrase)
	var input string
	fmt.Scanln(&input)
	return input
}
