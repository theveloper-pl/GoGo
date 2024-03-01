package main

import (
	"fmt"
	"sync"
)

var msg string
func updateMessage(s string, wg *sync.WaitGroup){
	defer wg.Done()
	msg = s
}

func main() {
	var messages = []string{"Eluwina","Siemanko","Co tam"}
	var wg sync.WaitGroup

	wg.Add(2)
	go updateMessage(messages[0], &wg)
	go updateMessage(messages[1], &wg)
	wg.Wait()
	
	fmt.Println(msg)


}
