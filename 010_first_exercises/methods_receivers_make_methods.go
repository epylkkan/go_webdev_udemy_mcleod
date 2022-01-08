package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

func (m person) speak() {
	fmt.Println("Hello, my name is", m.fName)
} 

func main() {
	// composite literal; struct literal
	p1 := person{"Todd", "McLeod"}
	p2 := person{"Nina", "Simone"}
	fmt.Println(p1)
	fmt.Println(p2)
	p1.speak()
	p2.speak()
}

// func (receiver) identifier(parameters) (returns) { <code> }
