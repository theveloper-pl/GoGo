package main

import (
	"fmt"
	"sync"
)

var msg string
func updateMessage(s string, wg *sync.WaitGroup, m *sync.Mutex){
	defer wg.Done()
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	var messages = []string{"Eluwina","Siemanko","Co tam"}
	var wg sync.WaitGroup

	var mutex sync.Mutex

	// for x := range messages{
	// 	wg.Add(1)
	// 	updateMessage(messages[x],&wg)
	// 	wg.Wait()
	// 	fmt.Println(msg)
	// }


	wg.Add(2)
	go updateMessage(messages[0], &wg, &mutex)
	go updateMessage(messages[1], &wg, &mutex)
	wg.Wait()
	
	fmt.Println(msg)


}
