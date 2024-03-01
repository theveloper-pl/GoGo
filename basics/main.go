package main

import (
	"fmt"
	// "rsc.io/quote"
)

func main1() {

    var name string = "Mateusz"
    name2 := "Adam"

    var a,b,c,d int = 1,32,5,2
    const PI = 3.14


	fmt.Println("What's your name ?")
	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(a,b,c,d)
	fmt.Println(PI)
    fmt.Printf("%v is type of %T",name,name)


    var arr1 = [3]int{1,2,3}
    arr2 := [5]int{4,5,6}

    var arr3 = [5]string{"adam","mati"}
  
    fmt.Println(arr1)
    fmt.Println(arr2)
    fmt.Printf("%#v",arr3)
}
