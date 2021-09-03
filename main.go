package main

/**
* This function will automatically called when we execute the project
 */
func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" -- Initialize and Automatically set the type of the variable with respect to the value assigned
	// cards := []string{newCard(), "Ace of Diamonds"}
	// cards := deck{newCard(), "Ace of Diamonds"} // replaced the slice of string with 'deck' type :- deck === []string
	cards := newDeck() // replaced the hard coded deck with newDeck()

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print() // replaced the for loop with the print function in deck.go

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.saveToFile("my_cards")

	// fmt.Println(cards)
}

func newCard() string {
	return "Five of Diamonds"
}
