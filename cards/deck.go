package main

import (
	"fmt"
	"strconv"
)

type Deck []string

func newDeck() Deck {
	d := Deck{}
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}

	for s := 0; s < len(suits); s++ {
		for c := 0; c < 13; c++ {
			var cardNum string
			switch c {
			case 0:
				cardNum = "Ace"
			case 10:
				cardNum = "Jack"
			case 11:
				cardNum = "Queen"
			case 12:
				cardNum = "King"
			default:
				cardNum = strconv.Itoa(c + 1)
			}

			cardName := cardNum + " of " + suits[s]
			d = append(d, cardName)
		}
	}
	return d
}

func (d Deck) deal(handSize int) Deck {
	hand := d[:handSize]
	d = d[handSize:]
	return hand
}

func (d Deck) print(label string) {
	fmt.Printf("\n%s (%d):\n %q", label, len(d), d)
}
