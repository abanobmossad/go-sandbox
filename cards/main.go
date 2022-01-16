package main

import "fmt"

func main() {
	deck := newDeck()
	hand := deck.deal(3)
	fmt.Printf("Hand %v", hand)
	deck.print()
}
