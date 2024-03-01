package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestConcurrencyk(t *testing.T){

	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup

	wg.Add(1)
	go printSmth("test", &wg)
	wg.Wait()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output,"test"){
		t.Errorf("Expected different value ")
	}

}

func TestPritning(t *testing.T){
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	printSmth2("test")
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut

	if !strings.Contains(output,"test"){
		t.Errorf("Expected different value xd ")
	}
}


