package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	cards,cards1 := deal(cards,5)
	fmt.Printf("%#v",cards)
	fmt.Printf("%#v",cards1)	
	writeToFile(cards, "deck.txt")
	var cards3 deck
	cards3 = readFromFile("deck.txt")
	fmt.Println("Shuffle cards :")
	cards3.print()
	shuffle(cards3)
	cards3.print()
	// cards := deck{newCard(),newCard()}
	// fmt.Printf("%#v \n", cards)
	// cards = append(cards,newCard())
	// fmt.Printf("%#v \n", cards)

	// cards.print()

}

