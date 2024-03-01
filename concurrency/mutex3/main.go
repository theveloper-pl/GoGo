package main

import (
	"fmt"
	"sync"
)



type Income struct{
	Source string
	Amount int
}

var bankBalance int
func main() {
	
	var wg sync.WaitGroup	
	var mutex sync.Mutex

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 2135},
		{Source: "Thefts", Amount: 43},
		{Source: "Second job", Amount: 523},
	}

	wg.Add(len(incomes))

	fmt.Printf("Initial account balance equal : $%d \n\n",bankBalance)	
	for i,incomes := range incomes{
		go doSomething(i,incomes,&wg,&mutex)		
	}
	wg.Wait()
	fmt.Printf("New back balance: $%d \n\n",bankBalance)
}

func doSomething(i int, income Income, wg *sync.WaitGroup, mutex *sync.Mutex){
	defer wg.Done()
	var incomeTotal Income
	incomeTotal.Source=income.Source
	for week:=1; week <=52; week+=1{
		incomeTotal.Amount += income.Amount
	}

	fmt.Printf("Source: %s, ammount: %d \n",incomeTotal.Source,incomeTotal.Amount)
	mutex.Lock()
	bankBalance+=incomeTotal.Amount
	mutex.Unlock()
}