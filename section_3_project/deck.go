package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*create a new type of "deck"
which is just a slice of strings*/
type deck []string

//it is good Go naming convention to have exported functions to start with an 
//uppercase letter, whereas functions that are used locally in a file start with
//a lowercase letter. Also organized these funcs on whether or not they have a
//reciever value or not.

func NewDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func Deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func NewDeckFromFile(fileName string) deck {
	byteSlice, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	strSlice := strings.Split(string(byteSlice), ",")

	return deck(strSlice)
}

func (d deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) SaveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}	

//I dont really like the instructors way of shuffling, it's kinda wack
//kinda unreadable too
//may refactor this later...
func (d deck) ShuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}