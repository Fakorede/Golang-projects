package main

import "fmt"

// create a new type of 'deck
// which is a slice of strings
type deck []string

// create deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// function with a reciever
// works like a method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal a hand of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
