package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

type contactInfo struct{
	email string
	phone int
}


func main() {
	p1 := person{firstName: "Mateusz", lastName: "Popielarski", contact: contactInfo{email: "mateusz268@gmail.com", phone: 997}}
	// fmt.Printf("%#v",p1)
	p1.readMe()
	p1.updateName("Adam")
	p1.readMe()
}

func (p person)readMe(){
	fmt.Printf("%#v",p)
}


func (p *person) updateName(newName string){
	(*p).firstName = newName
}