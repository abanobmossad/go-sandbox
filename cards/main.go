package main

func main() {
	deck := newDeck()
	hand := deck.deal(3)
	deck.print("Before  Deck")
	hand.print("Hand")
	deck.print("Deck")
}
