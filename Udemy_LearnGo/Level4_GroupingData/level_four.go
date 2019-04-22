// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// Author: Evan Douglass
// Created: April 21 2019

package main

import (
	"fmt"
)

// Arrays
func ex1() {
	var intArr [5]int
	for i := range intArr {
		intArr[i] = i + 1
	}
	fmt.Printf("%v\n%T\n", intArr, intArr)
}

// Slices
func ex2() {
	var intSli = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v\n%T\n", intSli, intSli)
}

// Slicing slices
func ex3() {
	var intSli = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intSli[:])
	fmt.Println(intSli[3:7])
	fmt.Println(intSli[:5])
	fmt.Println(intSli[5:])
	fmt.Println(intSli)
}

// More working with slices
func ex4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	fmt.Println("append")
	x = append(x, 52)
	fmt.Println(x)

	fmt.Println("and again...")
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	fmt.Println("append from another slice")
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

// Delete using append
func ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Println(y)
}

// Slice Properties
func ex6() {
	states := []string{
		` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`,
	}
	fmt.Println("len:", len(states))
	fmt.Println("cap:", cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}

// 2D slices
func ex7() {
	bond := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i := 0; i < len(bond); i++ {
		fmt.Println(bond[i])
		for j := 0; j < len(bond[i]); j++ {
			fmt.Print(bond[i][j], " ")
		}
		fmt.Println()
	}
}

// Maps
func ex8() {
	m := make(map[string][]string)
	m["bond_james"] = []string{
		"Shaken, not stirred",
		"Martinis",
		"Women",
	}

	m["moneypenny_miss"] = []string{
		"James Bond",
		"Literature",
		"Computer Science",
	}

	m["no_dr"] = []string{
		"Being evil",
		"ice cream",
		"Sunsets",
	}

	// Print by value
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println()

	delete(m, "bond_james")
	for k, v := range m {
		fmt.Println(k, v)
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
	ex6()
	fmt.Println()

	fmt.Println("Exercise 7")
	ex7()
	fmt.Println()

	fmt.Println("Exercise 8")
	ex8()
	fmt.Println()
}
