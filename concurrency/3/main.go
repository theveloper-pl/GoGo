package main

import (
	"fmt"
	"time"
)



func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go xd(c1)
	go xd2(c2)

	// for msg := range c{
	// 	fmt.Print(msg)	
	// }


	for {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
	}

}

func xd(c chan <-string){
	for {
	c <- "Elo"
	time.Sleep(time.Millisecond * 2000)		
	}
}


func xd2(c chan <-string){
	for {
	c <- "Siemanko"
	time.Sleep(time.Millisecond * 500)		
	}
}