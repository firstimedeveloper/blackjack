package blackjack

import (
	"fmt"
	"os"
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
	gameDeck := deck.New(deck.MultipleDecks(3), deck.Shuffle)

	GetInput("Press enter to play")

	players := make([]Player, numOfPlayers+1) //+1 as we also need a dealer
	var dealtCard deck.Card
	for round := 0; round < 2; round++ {
		for i := 0; i < numOfPlayers+1; i++ { //=1 because we need to account for the dealer
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
	if dealer.isBlackJack() {
		fmt.Println("Dealer got a natural blackjack! Game over.")
		os.Exit(1)
	}
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
			if players[i].isBlackJack() {
				fmt.Printf("Player %d got a blackjack!\n", i+1)
				endOfTurn = true
				continue
			}
			validInput := false
			var choice string
			for !validInput {
				fmt.Println("==============")
				fmt.Printf("Player %d\n", i+1)
				choice = GetInput("hit or stand? ")
				if choice == "hit" || choice == "stand" {
					validInput = true
				} else {
					fmt.Println("Enter a valid word")
				}

				switch choice {
				case "hit":
					//dealtCard = gameDeck[len(gameDeck)-1]
					//gameDeck = gameDeck[:len(gameDeck)-1]
					dealtCard, gameDeck = draw(gameDeck)

					//append dealtcard to hand
					players[i] = append(players[i], dealtCard)
					fmt.Printf("Player %d's hand: %s\n", i+1, players[i])
					fmt.Printf("Score: %d\n", players[i].getValueHand())
					if players[i].getValueHand() > 21 {
						fmt.Println("Busted")
						endOfTurn = true
					}
				case "stand":
					endOfTurn = true
					//do nothing
				}
			}
		}
	}
	fmt.Printf("Dealer's hand is: %s\nScore: %d\n", dealer, dealer.getValueHand())
	for dealer.getValueHand() < 17 || dealer.isSoftSeventeen() {
		//dealtCard = gameDeck[len(gameDeck)-1]
		//gameDeck = gameDeck[:len(gameDeck)-1]
		dealtCard, gameDeck = draw(gameDeck)
		dealer = append(dealer, dealtCard)
		fmt.Println("The dealer will hit")
		fmt.Printf("Dealer's new hand is: %s\nScore: %d\n", dealer, dealer.getValueHand())
	}
	if dealer.getValueHand() > 21 {
		fmt.Println("Dealer Busted!")
	} else {
		fmt.Println("The dealer will stand")

	}

	for i, player := range players {
		if player.getValueHand() < 21 {
			fmt.Println("==========")
			fmt.Printf("Player %d:\n", i+1)
			fmt.Printf("Score: %d\n", player.getValueHand())
			if dealer.getValueHand() > 21 {
				fmt.Println("Dealer busted so you win!")
			} else {
				fmt.Printf("Dealer score: %d\n", dealer.getValueHand())
				if player.getValueHand() > dealer.getValueHand() {
					fmt.Println("Better hand than the dealer, you win!")
				} else if player.getValueHand() == dealer.getValueHand() {
					fmt.Println("Same score as the dealer. Draw!")
				} else {
					fmt.Println("Dealer has a better hand.")
				}
			}
		}
	}
}

//draw draws the last card from the deck (dealtCard) and returns the dealtCard
//and the new deck without the dealt card.
func draw(d []deck.Card) (deck.Card, []deck.Card) {
	dealtCard := d[len(d)-1]
	d = d[:len(d)-1]
	return dealtCard, d
}

//isBlackJack would only return the correct value when the player has 2 cards
func (p Player) isBlackJack() bool {
	for _, card := range p {
		if card.Rank == deck.Ace && p.getValueHand() == 21 {
			return true
		}
	}
	return false
}

//isSoftSeventeen returns true if the player has an ace
//used to determine if the dealer has a soft 17
func (p Player) isSoftSeventeen() bool {

	for _, card := range p {
		if card.Rank == deck.Ace && p.getValueHand() == 17 {
			return true
		}
	}
	return false
}

func (p Player) getValueHand() int {
	//I noticed that the number of different combinations of a hand's value with aces
	//is 1 + # of aces.
	//Example
	//A=Ace O=One (where the value of the Ace card can be 11 or 1)
	//if there's one Ace card in a hand, the possiblility is:
	//A, O -- 2 possibilities
	//Two Ace cards:
	//OO, AO, OO -- 3 possibilities (OA doesn't count as order doesn't change the score)

	var score int
	var aces int
	for _, c := range p {
		if c.Rank == deck.Ace {
			aces++ //counts the number of aces
		}
		score += getValueCard(c)
	}
	if aces == 0 {
		return score //the score isn't variable if the hand doesn't contain aces.
	}
	scores := make([]int, aces+1) //A slice of ints with the capacity of # of aces+1

	count := 0
	for i := 0; i < cap(scores); i++ {
		scores[i] += score + (cap(scores)-count)*10 //This will order the slice in descending order (highest to lowest)
		count++
	}

	for _, s := range scores {
		if s <= 21 { //since the order of the slice is descending, the first number below 21 is the score
			score = s
		}
	}

	return score
}

func getValueCard(c deck.Card) int {
	//returns the integer score of a specific card
	//Ace equals to one
	switch c.Rank {
	case deck.Jack, deck.Queen, deck.King:
		return 10
	default:
		return int(c.Rank)
	}
}

//not sure if this is a good idea or not
//but created the function initially cause I didn't want to keep writing the combination.
func GetInput(phrase string) string {
	fmt.Print(phrase)
	var input string
	fmt.Scanln(&input)
	return input
}
