package main

import (
	"fmt"
	"strconv"
)

type Deck []string

func newDeck() Deck {
	d := Deck{}
	suits := []string{"Spade", "Heart", "Diamond", "Club"}

	for s := 1; s < len(suits); s++ {
		for c := 1; c < 13; c++ {
			var cardNum string
			switch c {
			case 1:
				cardNum = "Ace"
			case 11:
				cardNum = "Jack"
			case 12:
				cardNum = "Queen"
			case 13:
				cardNum = "King"
			default:
				cardNum = strconv.Itoa(c)
			}

			cardName := cardNum + " of " + suits[s]
			d = append(d, cardName)
		}
	}
	return d
}

func (d Deck) deal(handSize int) []string {
	return d[:handSize]
}

func (d Deck) print() {
	fmt.Printf("Deck cards:\n %q", d)
}
