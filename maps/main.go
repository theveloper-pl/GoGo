package main

import (
	"fmt"
)


func main() {

	// colors2 := make(map[int]string)
	// delete(colors2, "blue")


	colors := map[string]string{
		"blue":"xdd",
		"adam":"saddsa",
	}

	fmt.Println("Siemka !")
	colors["yellow"] = "jeszcze jok"
	fmt.Printf("%v", colors["blue"])
	fmt.Printf("%v", colors)

	kolory(colors)

}

func kolory(c map[string]string){
	for i,v := range c{
		fmt.Printf("%v : %v \n",i,v)
	}
}