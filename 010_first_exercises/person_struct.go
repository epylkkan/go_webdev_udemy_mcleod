package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type sa struct {
	person
	ltk bool
}

func (p person) pSpeak() {
	fmt.Println(p.fname, `says, "Wassup!"`)
}

func (s sa) saSpeak() {
	fmt.Println(s.fname, "has a license to kill:", s.ltk)
}

func main() {
	p := person{"Todd", "McLeod"}
	s := sa{person{"James", "Bond"}, true}
	fmt.Println(p.fname)
	p.pSpeak()
	fmt.Println(s.fname)
	s.saSpeak()
	s.pSpeak()
}