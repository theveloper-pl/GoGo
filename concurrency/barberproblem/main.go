package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 2
var arrivalRate = 100
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second


func main(){

	rand.Seed(time.Now().UnixNano())

	color.Yellow("The Sleeping Barber Problem")
	color.Yellow("---------------------------")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	barberShop := BarberShop{
		ShopCapacity: seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		BarberDone: doneChan,
		ClientsChan: clientChan,
		Open: true,
	}

	color.Green("Shop is open for today !")

	barberShop.addBarder("Frank")
	barberShop.addBarder("Adam")
	barberShop.addBarder("Sebastian")
	barberShop.addBarder("Krzysiek")
	barberShop.addBarder("Maciek")
	barberShop.addBarder("Marian")

	shopClosing := make(chan bool)
	closed := make(chan bool)

	go func(){

		<-time.After(timeOpen)
		shopClosing <- true
		barberShop.closeShop()
		closed <- true
	}()

	i:= 1
	go func(){
		for{
			randomMilisencods := rand.Int() % (2 * arrivalRate)
			select {
			case <- shopClosing:
				return
			case <- time.After(time.Millisecond * time.Duration(randomMilisencods)):
				barberShop.addClient(fmt.Sprintf("Client #%d",i))
				i++
			}
		}
	}()

		<-closed
}