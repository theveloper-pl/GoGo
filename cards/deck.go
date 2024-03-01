package main

import (
	 "fmt"
	 "log"
	 "io/ioutil"
	 "strings"
	 "math/rand"
	 "time"
)

type deck []string

func (cards deck) print() {
	fmt.Println("\nThis deck :")
	for i, card := range cards{
		fmt.Println(i,card)
	}
}

func deal(cards deck, amount int) (deck, deck){
	return cards[:amount], cards[amount:]
}

func newDeck() deck{
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	var cards deck

	for _,suit:= range cardSuits{
		for _,value:= range cardValues{
			cards = append(cards,suit+" of "+value)
		}
	}
	return cards
}

func newCard() string{
	return "New card"
}


func readFromFile(filename string) deck{
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return deck(strings.Split(string(content),";"))
}


func writeToFile(d deck, filename string){
	message := []byte(strings.Join([]string(d), ";"))
	err := ioutil.WriteFile(filename, message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func shuffle (d deck) deck{
    rand.Seed(time.Now().UnixNano())
    shuffledDeck := []string(d)
    rand.Shuffle(len(shuffledDeck), func(i, j int) {
        shuffledDeck[i], shuffledDeck[j] = shuffledDeck[j], shuffledDeck[i]
    })
    return shuffledDeck
}

