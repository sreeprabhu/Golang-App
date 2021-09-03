package main

import "fmt"

/**
* Create a new type of 'deck'
* which is a slice of strings
 */
type deck []string

/**
* newDeck() - Creates a new deck
 */
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Boots"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

/**
* print() - Logs the contents of a deck of cards
* (d deck) - receiver of function print()
* any variable of type deck now gets access to the function print()
 */
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
