// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// Author: Evan Douglass
// Created: April 21 2019

package main

import (
	"fmt"
)

// Basic for loop
func ex1() {
	for i := 1; i <= 10000; i++ {
		// using this format to save output space
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Nested loops
func ex2() {
	for i := 65; i < 91; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\n", i)
		}
	}
}

// for/while loop
func ex3() {
	year := 1990
	for year < 2020 {
		fmt.Print(year, " ")
		year++
	}
	fmt.Println()
}

// for with break
func ex4() {
	year := 1990
	for {
		if year == 2020 {
			fmt.Println()
			break
		}
		fmt.Print(year, " ")
		year++
	}
}

// Using modulo
func ex5() {
	for i := 10; i < 101; i++ {
		mod := i % 4
		fmt.Print(mod, " ")
	}
	fmt.Println()
}

// Using switch
func ex8() {
	switch {
	case false:
		fmt.Println("Won't run")
	case 1 == 1:
		fmt.Println("Hello, from switch!")
	}
}

// More switch
func ex9() {
	favSport := "Lawn Darts"
	switch favSport {
	case "Football":
		fmt.Println("American or World?")
	case "Soccer":
		fmt.Println("That is more clear")
	case "Lawn Darts":
		fmt.Println("Really???")
	default:
		fmt.Println("I guess there are no other sports...")
	}
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
	fmt.Println("Skipping ex6, just an if statement")
	fmt.Println()

	fmt.Println("Exercise 7")
	fmt.Println("Skipping ex6, just an if-else statement")
	fmt.Println()

	fmt.Println("Exercise 8")
	ex8()
	fmt.Println()

	fmt.Println("Exercise 9")
	ex9()
}
