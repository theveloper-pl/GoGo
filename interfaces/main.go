package main

import (
	"fmt"
)

type bot interface{
	getGreeting() string
	// getGreeting() (string, int) (string, error)
}

type englishBot struct{
	greeting string
	language string
}

type spanishBot struct{
	greeting string
	language string
}




func main() {

	s1 := englishBot{greeting:"Siemka", language:"English"}
	s2 := spanishBot{greeting:"Ola soi dora !", language:"Spanish"}

	printGreeting(s1)
	printGreeting(s2)

}

func printGreeting(b bot){
	fmt.Printf("%#v\n",b.getGreeting())
}

func (eb englishBot) getGreeting() string{
	return "Hello there"
}

func (sb spanishBot) getGreeting() string{
	return "Ola soi dora !"
}
