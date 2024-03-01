package main

import (
	"fmt"
	"sync"
	"net"
)

func main() {

	addrs, err := net.LookupHost("theveloper.pl")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(addrs)


	// var wg sync.WaitGroup

	// words := []string{"jeden","dwa","trzy","cztery"}

	// wg.Add(len(words))	
	// for i,e := range words{
	// 	go printSmth(fmt.Sprintf("%d %s",i,e), &wg)
	// }
	// wg.Wait()


	// wg.Add(1)
	// printSmth("dwa", &wg)
}


func printSmth(s string, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(s)
	
}

func printSmth2(s string){
	fmt.Println(s)
	
}
