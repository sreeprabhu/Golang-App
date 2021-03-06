package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log the error and exit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano()) // check the rand documentation
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)           // generate the random number
		d[i], d[newPosition] = d[newPosition], d[i] // swap the elements
	}
}
