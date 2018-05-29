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
		fmt.Printf("player %d's hand: %s\n", i+1, player)
	}
	fmt.Printf("Dealer's hand: %s\n", dealer.DealerString())

	//testing purposes
	//printing the scores
	for i, player := range players {
		fmt.Printf("Player %d's current score: %d\n", i+1, player.getValueHand())
	}

	for i := range players {
		endOfTurn := false
		for !endOfTurn {
			if players[i].getValueHand() == 21 {
				fmt.Printf("Player %d got a blackjack!", i+1)
				endOfTurn = true
				continue
			}
			validInput := false
			var choice string
			for !validInput {
				fmt.Println("==============")
				fmt.Printf("Player %d\n", i+1)
				choice = getInput("hit or stand? ")
				if choice == "hit" || choice == "stand" {
					validInput = true
				} else {
					fmt.Println("Enter a valid word")
				}

				switch choice {
				case "hit":
					dealtCard = gameDeck[len(gameDeck)-1]
					gameDeck = gameDeck[:len(gameDeck)-1]
					//append dealtcard to hand
					players[i] = append(players[i], dealtCard)
					fmt.Printf("Player %d's hand: %s\n", i+1, players[i])
					fmt.Printf("Score: %d\n", players[i].getValueHand())
					if players[i].getValueHand() > 21 {
						fmt.Println("Busted")
						endOfTurn = true
					}
				case "stand":
					if players[i].getValueHand() > dealer.getValueHand() {
						fmt.Printf("Dealer's hand is: %s\nScore: %d\n\n", dealer, dealer.getValueHand())
						fmt.Println("Better hand than the dealer, you win!")

					} else {
						fmt.Printf("Dealer's hand is: %s\nScore: %d\n", dealer, dealer.getValueHand())
						fmt.Println("Dealer has a better hand.")
					}
					endOfTurn = true
					//do nothing
				}
			}
		}
	}
}

func (p Player) getValueHand() int {
	var score int
	var aces int
	for _, c := range p {
		if c.Rank == deck.Ace {
			aces++
		}
		score += getValueCard(c)
	}
	if aces == 0 {
		return score
	}
	scores := make([]int, aces+1)

	count := 0
	for i := 0; i < cap(scores); i++ {
		scores[i] += score + (cap(scores)-count)*10
		count++
	}

	for _, s := range scores {
		if s <= 21 {
			score = s
		}
	}

	return score
}

func getValueCard(c deck.Card) int {
	switch c.Rank {
	case deck.Jack, deck.Queen, deck.King:
		return 10
	default:
		return int(c.Rank)
	}
}

func getInput(phrase string) string {
	fmt.Print(phrase)
	var input string
	fmt.Scanln(&input)
	return input
}
