// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// Author: Evan Douglass
// Created: April 21 2019

package main

import (
	"fmt"
)

// Problem 1
func num1() {
	// Use short declaration operator to assign different types
	x := 42
	y := "James Bond"
	z := true

	// Print all in one statement
	fmt.Println(x, y, z)

	// Print separately
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

// Problem 2
// Declare vars in package scope
var x int
var y string
var z bool

func num2() {
	fmt.Println("Printing global vars with zero values.")
	fmt.Println(x, y, z)
}

// Problem 3
func num3() {
	// Assign global variables
	x = 43
	y = "James Bend"
	z = false

	// Sprintf into string var
	s := fmt.Sprintf("x: %v, y: %v, z: %v", x, y, z)
	fmt.Println(s)
}

// Problem 4
// Create type
type ballin int

// Var for new type
var b ballin

func num4() {
	// Print zero value and type
	fmt.Printf("b: %v, type of x: %T\n", x, x)
	// Assign value and print that
	b = 75
	fmt.Println("b is now", b)
}

// Problem 5
// Make var for underlying type of ballin
var a int

func num5() {
	a = int(b)
	fmt.Println("a:", a)
}

func main() {
	fmt.Println("Exercise 1")
	num1()
	fmt.Println()

	fmt.Println("Exercise 2")
	num2()
	fmt.Println()

	fmt.Println("Exercise 3")
	num3()
	fmt.Println()

	fmt.Println("Exercise 4")
	num4()
	fmt.Println()

	fmt.Println("Exercise 5")
	num5()
	fmt.Println()
}
