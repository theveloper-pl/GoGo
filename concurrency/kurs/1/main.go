package main

import (
	"fmt"
)

const _ = iota

func main() {

	x := []byte("siemanko")
	fmt.Println(string(x[0]))

	a := 1
	fmt.Printf("%v  %b",a,a)

	for  i := 0; i<=10; i++ {
		a = a << i
		fmt.Printf("%v  %b\n",a,a)
	}



}
