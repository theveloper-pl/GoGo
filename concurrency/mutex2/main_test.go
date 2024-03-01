package main

import (
	"strings"
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T){

	var wg sync.WaitGroup

	wg.Add(1)
	updateMessage("newMessage",&wg)
	wg.Wait()

	if !strings.EqualFold(msg,"newMessage"){
		t.Errorf("Wrong message !")
	}

}