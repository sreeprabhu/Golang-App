package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

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

/**
* deal() - returns two sets of decks
* d[:handSize] - return the set of elements starting from index 0 to the handSize count
* d[handSize:] - return the set of elements until the last index after the handSize count of elements
 */
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

/**
* toString() receiver function
* converts sllice of string (deck of cards) to string
* using inbuild func join(str []string, seperator string) string
 */
func (d deck) toString() string {
	return strings.Join([]string(d), ",") // type casted the deck to []string
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666) // any one can read and write this file
}
