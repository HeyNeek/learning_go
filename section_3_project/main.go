package main

func main() {
	cards := newDeckFromFile("my_cards")

	cards.print()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}