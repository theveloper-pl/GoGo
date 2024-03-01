package main

import (
	"github.com/fatih/color"
	"time"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarberDone chan bool
	ClientsChan chan string
	Open bool
}

func (b *BarberShop) closeShop(){
	color.Cyan("Closing Barber shop for today !")
	close((*b).ClientsChan)
	(*b).Open = false

	for a:= 1; a<=(*b).NumberOfBarbers; a++{
		<-(*b).BarberDone
	}
	close((*b).BarberDone)

	color.Green("-------------------------")
	color.Green("The barbershop is now closed for the day, and everyone has gone home.")
}

func (b *BarberShop) addBarder(barber string){

	(*b).NumberOfBarbers+=1

	go func(){
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clienrs", barber)

		for {
			if len((*b).ClientsChan) == 0{
				color.Yellow("There is nothing to do so %s takes a nap", barber)
				isSleeping = true
			}

			client, shopOpen := <-(*b).ClientsChan
			if shopOpen{
				if isSleeping{
					color.Yellow("%s wakes %s", client, barber)
				}

				b.cutHair(barber,client)

			}else{
				b.sentBarberHome(barber)
				return
			}
		}

	}()

}

func (b *BarberShop) cutHair(barber string, client string){
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep((*b).HairCutDuration)
	color.Green("%s is done cutting %s's hair", barber, client)
}

func (b *BarberShop) sentBarberHome(barber string){
	color.Cyan("%s is going home", barber)
	(*b).BarberDone <- true
}

func (b *BarberShop) addClient(client string){
	color.Green("*** %s arrives !!!", client)

	if (*b).Open{

		select{
		case (*b).ClientsChan <- client:
				color.Blue("%s takes a seat in the waiting room", client)
			
		default:
			color.Red("The waiting room is full, so %s leaves", client)

			
		}

			
	}else{
		color.Red("The shop is already closed, so %s leaves!", client)
	}

	
}