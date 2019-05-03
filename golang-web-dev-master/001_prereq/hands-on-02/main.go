package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	location string
}

func (p person) pSpeak(say string) {
	fmt.Println(say)
}

func (s secretAgent) saSpeak(say string) {
	fmt.Printf("%s, darling.\n", say)
}

func main() {
	p := person{"Evan", "Douglass", 28}
	sa := secretAgent{
		person{"James", "Bond", 34},
		"London, UK",
	}

	fmt.Println(p.first)
	p.pSpeak("Hello!")

	fmt.Println(sa.first)
	sa.saSpeak("Hello!")

	sa.pSpeak("Hello!")
}
