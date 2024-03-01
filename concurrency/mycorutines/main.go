package main

import (
    "fmt"
	"math/rand"
	"time"
)

type Communicator struct{
    dataInt chan int
    dataString chan string
}

func goCrazy(c Communicator){
    rand.Seed(time.Now().UnixNano())
    time := rand.Intn(5)
    if time < 3{
        c.dataInt <- time
    }else{
        c.dataString <- fmt.Sprintf("Time is equal %d",time)
    }

}



func main() {
    c := Communicator{dataInt: make(chan int), dataString: make(chan string)}

    go goCrazy(c)

    for i := 0; i < 1; i++ {
        select {
        case msg1 := <-c.dataInt:
            fmt.Println("received", msg1)

        case msg2 := <-c.dataString:
            fmt.Println("Nottt", msg2)
        }
    }
}