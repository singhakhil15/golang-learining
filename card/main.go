package main

func main() {
	//cards := newDeck()

	//cards.print()
	//fmt.Println(cards)

	// hands, remainingDeck := deal(cards, 5)
	// fmt.Println(hands)
	// fmt.Println(remainingDeck)

	// hands.print()
	// handsB, remainingDeck := deal(remainingDeck, 5)

	// handsB.print()
	// fmt.Println(" m mmmmmmmmmmmmmmmmmmmmmmm  b  bbbbbbbbbbbbbbbbbbb ")
	// fmt.Println(" m mmmmmmmmmmmmmmmmmmmmmmm  b  bbbbbbbbbbbbbbbbbbb ")
	// remainingDeck.print()

	// cards.saveToFile("my_cards")
	// hands.saveToFile("my_hands")
	// remainingDeck.saveToFile("remaining_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
