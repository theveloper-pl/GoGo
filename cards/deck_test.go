package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T){
	cards := newDeck()
	realCards := newDeck()

    if len(cards) != len(realCards) {
        t.Errorf("Lenght of decks does not equal predicted value !")
    }

    for i, v := range cards {
        if v != realCards[i] {
            t.Errorf("Created Deck does not match predicted output !")
        }
    }

}


func TestDealDeck(t *testing.T){
	cards := newDeck()
	expected1 := deck{"Spades of Ace", "Spades of Two", "Spades of Three", "Spades of Four", "Hearts of Ace"}
	expected2 := deck{"Hearts of Two", "Hearts of Three", "Hearts of Four", "Diamonds of Ace", "Diamonds of Two", "Diamonds of Three", "Diamonds of Four", "Clubs of Ace", "Clubs of Two", "Clubs of Three", "Clubs of Four"}
	result1, result2 := deal(cards,5)

    if len(expected1) != len(result1) {
        t.Errorf("Lenght of decks does not equal predicted value !")
    }

    for i, v := range expected1 {
        if v != result1[i] {
            t.Errorf("Created Deck does not match predicted output !")
        }
    }


    if len(expected2) != len(result2) {
        t.Errorf("Lenght of decks does not equal predicted value !")
    }

    for i, v := range expected2 {
        if v != result2[i] {
            t.Errorf("Created Deck does not match predicted output !")
        }
    }

}

func TestFile(t *testing.T){
	cards := newDeck()
	writeToFile(cards, "_testFile.txt")
	cardsFromFile := deck(readFromFile("_testFile.txt"))

    if len(cards) != len(cardsFromFile) {
        t.Errorf("Lenght of deck does not equal predicted value !")
    }

    for i, v := range cards {
        if v != cardsFromFile[i] {
            t.Errorf("Element %v does not match %v", v, cardsFromFile[i])
        }
    }
	os.Remove("_testFile.txt")

}