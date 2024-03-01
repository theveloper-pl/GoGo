package main

import (
	"fmt"
	"sync"
)


func printsmth(wg *sync.WaitGroup, name string){
	defer wg.Done()
	fmt.Printf("%s is done\n", name)	
}

func main() {

	names := []string{"jeden","dwa","trzy","cztery"}


	wg := &sync.WaitGroup{}
	wg.Add(len(names))


	for i := 0; i<len(names); i++{
		go printsmth(wg, names[i])		
	}
	fmt.Println("xdd")
	wg.Wait()
	fmt.Println("Koniec")
}