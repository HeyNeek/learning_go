package main

import "fmt"

func main() {
	fmt.Println("Reading Deck from file... ")
	cards := NewDeckFromFile("my_cards")
	fmt.Println("")

	fmt.Println("Read successful. Current cards: ")
	cards.Print()
	fmt.Println("")

	fmt.Println("Shuffling Cards: ")
	cards.ShuffleDeck()
	fmt.Println("")

	fmt.Println("After Shuffling: ")
	cards.Print()
	fmt.Println("")

	fmt.Println("Dealing cards...")
	hand, remainingCards := Deal(cards, 5)
	fmt.Println("Cards have been dealt. Inspect your hand and the remaining cards.")
	fmt.Println("")

	fmt.Println("Hand: ")
	hand.Print()
	fmt.Println("")

	fmt.Println("Remaining Cards: ")
	remainingCards.Print()
	fmt.Println("")
}