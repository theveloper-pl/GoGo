package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	websites := []string{"http://google.com","http://facebook.com","http://stackoverflow.com","http://golang.org","http://amazon.com"}

	c := make(chan string)

	for _, value := range websites{
		go ping(value, c)
		
	}
	for l := range c{
		go func(link string){
			time.Sleep(5 * time.Second)
			ping(link,c)
		}(l)
	}

}

func ping(domain string, c chan string) (int,error){
	_, err := http.Get(domain)
	if err != nil{
		// c <- domain+" is down ?"
		fmt.Printf("%v is Down ! \n",domain)
		c <- domain
		return 0, err
	}
	// c <- domain+" is up ?"
	fmt.Printf("%v is Up ! \n",domain)
	c <- domain
	return 1, err
}