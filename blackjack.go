package blackjack

import (
	"fmt"
	"strings"

	"github.com/firstimedeveloper/deck"
)

//Player type of slice of cards
type Player []deck.Card

//String function returns the player's hand in a formated string
//taken from the gophercises cource
func (p Player) String() string {
	str := make([]string, len(p))
	for i := range p {
		str[i] = p[i].String()
	}
	return strings.Join(str, ", ")
}

//DealerString function returns a formated string for the dealer
//taken from gophercises
func (p Player) DealerString() string {
	return p[0].String() + ", Hidden"
}

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
	dealer := players[len(players)-1]  //dealer is the last player in the players slice
	players = players[:len(players)-1] //the new players slice doesn't contain the dealer
	//printing dealt cards, while not showing last card
	//seems unnecessarily complicated
	//TODO make this shorter or something.
	for i, player := range players {
		fmt.Printf("player %d: %s\n", i+1, player)
	}
	fmt.Printf("Dealer: %s\n", dealer.DealerString())

	//testing purposes
	//printing the scores
	for i, player := range players {
		fmt.Printf("Player %d score: %d\n", i+1, player.getValueHand())
	}

	for _, player := range players {
		endOfTurn := false
		for !endOfTurn {
			validInput := false
			var choice string
			for !validInput {
				choice = getInput("hit or stand?")
				if choice == "hit" || choice == "stand" {
					validInput = true
				}
			}
			switch choice {
			case "hit":
				dealtCard = gameDeck[len(gameDeck)-1]
				gameDeck = gameDeck[:len(gameDeck)-1]
				//append dealtcard to hand
				player = append(player, dealtCard)
				fmt.Println(player)
			case "stand":
				endOfTurn = true
				//do nothing
			}

		}
	}

	for i, player := range players {
		fmt.Printf("Player %d score: %d\n", i+1, player.getValueHand())
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

func (p Player) getValueHand() int {
	var score int
	for _, c := range p {
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
