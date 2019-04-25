// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on pointers.
// Author: Evan Douglass
// Created: April 24 2019

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func dontChangeMe(p person) {
	p.first = "Mr."
	p.last = "T"
	p.age = 1000
}

func changeMe(p *person) {
	(*p).first = "Mr."
	(*p).last = "T"
	(*p).age = 1000
}

func main() {
	// Ex. 1: Print a pointer
	x := 10
	fmt.Printf("Address of x: %p\n", &x)
	// Ex. 2
	p := person{"Evan", "Douglass", 28}
	fmt.Println(p)
	dontChangeMe(p)
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
