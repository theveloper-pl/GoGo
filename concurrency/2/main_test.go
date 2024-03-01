package main

import (
	// "fmt"
	// "io"
	// "os"
	"strings"
	"io"
	"os"
	"sync"
	"testing"
)


func TestUpdateMessage(t *testing.T){
	var wg sync.WaitGroup
	wg.Add(1)
	go updateMessage("siemka", &wg)
	wg.Wait()
	if "siemka" != msg{
		t.Errorf("Values are not equal")
	}
}


func TestPrintMessage(t *testing.T){
	msg = "Testing message !"
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printMessage()

	_ = w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output, msg){
		t.Errorf("Expected: %s, Recived: %s",msg, output)
	}

}