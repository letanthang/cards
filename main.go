package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	cards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()

}
