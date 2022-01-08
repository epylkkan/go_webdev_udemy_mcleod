package main

import (
	"fmt"
)

type dog int

func (e dog) barkNum() {
	fmt.Println("I'm a dog", e)
}

type person struct {
	fName string
	lName string
}

func (m person) speak() {
	fmt.Println("Hello, my name is", m.fName)
} 

func main() {
	var d dog
	d = 17
	fmt.Printf("%T\n", d)
	d.barkNum()
	
	var g dog
	g = 42
	fmt.Printf("%T\n", g)
	g.barkNum()
	
	// composite literal; struct literal
	p1 := person{"Todd", "McLeod"}
	p2 := person{"Nina", "Simone"}
	fmt.Println(p1)
	fmt.Println(p2)
	p1.speak()
	p2.speak()
}

// func (receiver) identifier(parameters) (returns) { <code> }