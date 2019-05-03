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

func (p person) speak(say string) {
	fmt.Println(say)
}

func (s secretAgent) speak(say string) {
	fmt.Printf("%s, darling.\n", say)
}

type speaker interface {
	speak(string)
}

func what(sp speaker, str string) {
	fmt.Println("I said,", str)
}

func main() {
	p := person{"Evan", "Douglass", 28}
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   34,
		},
		location: "London",
	}

	what(p, "Hu?")
	what(sa, "Who has the codes?!?")
}
