package main

import (
	"testing"
	"sync"
)

func TestUpdateMessage(t *testing.T){
	var wg sync.WaitGroup	
	var mutex sync.Mutex

	incomesCheck := []Income{
		{Source: "Main job", Amount: 5230},
		{Source: "Gifts", Amount: 23},
		{Source: "Thefts", Amount: 4541},
		{Source: "Second job", Amount: 2315},
	}
	expectedBalance := 629668

	wg.Add(len(incomesCheck))
	for i,incomes := range incomesCheck{
		go doSomething(i,incomes,&wg,&mutex)		
	}
	wg.Wait()

	if bankBalance != expectedBalance{
		t.Errorf("Expected: %d", bankBalance)	
	}
	
}


func TestDoSomething(t *testing.T){
	var wg sync.WaitGroup	
	var mutex sync.Mutex	
	income := Income{Source: "Main job", Amount: 5230}
	bankBalance = 0
	expectedBalance := 5230 * 52
	wg.Add(1)
	doSomething(1,income,&wg,&mutex)
	wg.Wait()

	if bankBalance!=expectedBalance{
		t.Errorf("Bank balance equal %d, expecterd %d", bankBalance, expectedBalance)
	}
}