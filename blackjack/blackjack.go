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
	gameOver := false
	for !gameOver {
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
		for i, player := range players {
			for j, card := range player {
				fmt.Printf("Player %d card %d: %v\n", i+1, j+1, card)
			}
		}

		gameOver = true
	}
}

func getValue(c deck.Card) int {
	switch c.Rank {
	case deck.Jack, deck.Queen, deck.King:
		return 10
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
