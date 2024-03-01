package main

import (
	"fmt"
	"sync"
	"time"
)


func main(){
	var wg sync.WaitGroup
	wg.Add(1)
	go shout(&wg)
	wg.Wait()
}

func shout(wg *sync.WaitGroup){
	for x:=0; x<10; x++ {
		fmt.Println("Executing loop")
		time.Sleep(time.Millisecond * 500)
	}
	wg.Done()
}