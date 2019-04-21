// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// Author: Evan Douglass
// Created: April 21 2019

package main

import (
	"fmt"
)

// Prints the number 42 in decimal, binary and hex
func ex1() {
	num := 42
	fmt.Printf("%d\t%b\t%0x\n", num, num, num)
}

// Prints several boolean values
func ex2() {
	a := true == false
	b := 39 <= 33
	c := 99 >= 50
	d := true != false
	e := 55 < 55
	f := 55 > 54
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

// Prints typed and untyped constants
func ex3() {
	const (
		a = 42
		b = "Hello"
		c = false
	)

	const (
		d int    = 43
		e string = "Goodby"
		f bool   = true
	)

	fmt.Printf("%d, %d\n%s, %s\n%t, %t\n",
		a, d, b, e, c, f)
}

// Practice shifting bits
func ex4() {
	num := 84
	fmt.Printf("%d, %b, %0x\n", num, num, num)

	// shift bits
	num <<= 1
	fmt.Printf("%d, %b, %0x\n", num, num, num)
}

// Using raw string literals
func ex5() {
	s := `Hello,
	this is   
a raw
	string`
	fmt.Println(s)
}

// Using iota
func ex6() {
	// constants for next 4 years
	const (
		_ = iota
		year1
		year2
		year3
		year4
	)

	fmt.Println(year1, year2, year3, year4)
}

func main() {
	fmt.Println("Exercise 1")
	ex1()
	fmt.Println()

	fmt.Println("Exercise 2")
	ex2()
	fmt.Println()

	fmt.Println("Exercise 3")
	ex3()
	fmt.Println()

	fmt.Println("Exercise 4")
	ex4()
	fmt.Println()

	fmt.Println("Exercise 5")
	ex5()
	fmt.Println()

	fmt.Println("Exercise 6")
	ex6()
}
