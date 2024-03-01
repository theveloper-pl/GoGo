package main

import (
	"math/rand"
	"fmt"
	"time"
	"github.com/fatih/color"
)


const numberOfPIzzas = 10
var pizzasMade, pizzasFailed, total int

type Producer struct{
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct{
	pizzaNumber int
	message string
	success bool
}

func pizzaria(producer *Producer){
	var i = 0

	for{
		currentPizza := makePizza(i)
		if currentPizza != nil{
			i = currentPizza.pizzaNumber

			select{
			case producer.data <- *currentPizza:

			case quitChan := <- producer.quit:
				close(producer.data)
				close(quitChan)
				return
			}


		}
	}
}

func (p *Producer) Close() error{
	ch := make(chan error)
	p.quit <- ch
	return <- ch
}

func makePizza(pizzaNumber int) *PizzaOrder{
	pizzaNumber++
	if pizzaNumber <= numberOfPIzzas{
		deley := rand.Intn(5) + 1
		fmt.Printf("Recived order #%d\n",pizzaNumber)

		rnd := rand.Intn(12)+1
		msg := ""
		success := false

		if rnd <5{
			pizzasFailed++
		}else{
			pizzasMade++
		}
		total++

		fmt.Printf("Making pizza no. %d, it will take %d seconds ... \n",pizzaNumber,deley)
		time.Sleep(time.Duration(deley)*time.Second)

		if rnd <=2{
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d",pizzaNumber)
		}else if rnd <=4{
			msg = fmt.Sprintf("*** The cook died while making pizza #%d",pizzaNumber)
		}else{
			msg = fmt.Sprintf("*** Succesfully made pizza #%d",pizzaNumber)
			success = true
		}


		currentPizza := PizzaOrder{pizzaNumber,msg, success}
		return &currentPizza
	}
	return &PizzaOrder{pizzaNumber: pizzaNumber}


}

func main() {
    rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(12))
	color.Cyan("|-----------------------------------|")
	color.Cyan("| The pizzeria is open for business |")
	color.Cyan("|-----------------------------------|")

	pizzaJob := Producer{data: make(chan PizzaOrder), quit: make(chan chan error),}
	go pizzaria(&pizzaJob)

	for i := range pizzaJob.data{
		if i.pizzaNumber <= numberOfPIzzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order number %d is out for delivery !",i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is really mad")
			}
		}else{
			color.Cyan("Done making pizzas ...")
			err := pizzaJob.Close()
			if err != nil{
				color.Red("Error closing channel", err)
			}
		}
	}

	color.Blue("We have made %d pizzas in total, %d was done well and %d was fucked up",total,pizzasMade,pizzasFailed)

}

