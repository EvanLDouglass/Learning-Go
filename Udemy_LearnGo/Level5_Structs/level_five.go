// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// Author: Evan Douglass
// Created: April 23 2019

package main

import "fmt"

type (
	person struct {
		first       string
		last        string
		favIceCream string
	}

	vehicle struct {
		doors int
		color string
	}

	truck struct {
		vehicle
		fourWheel bool
	}

	sedan struct {
		vehicle
		luxury bool
	}
)

var (
	p1 = person{
		first:       "Evan",
		last:        "Douglass",
		favIceCream: "Half Baked",
	}
	p2 = person{
		first:       "Santa",
		last:        "Claus",
		favIceCream: "Vanilla?",
	}
)

func ex1() {
	fmt.Println(p1)
	fmt.Println(p2.first, p2.last, p2.favIceCream)
}

func ex2() {
	m := make(map[string]person)
	m[p1.last] = p1
	m[p2.last] = p2
	for k, v := range m {
		fmt.Println(k, v)
	}
}

// Using embedded structs
func ex3() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println("s.doors:", s.doors, "s.luxury:", s.luxury)
	fmt.Println("t.doors:", t.doors, "t.fourWheel:", t.fourWheel)
}

// Using anonymous structs
func ex4() {
	anon := struct {
		one string
		two string
		s   struct {
			three string
		}
	}{
		one: "hello",
		two: "goodby",
		s: struct {
			three string
		}{
			three: "Anonymous & Embedded!",
		},
	}
	fmt.Println(anon)
}

func main() {
	fmt.Println("Example 1")
	ex1()
	fmt.Println()

	fmt.Println("Example 2")
	ex2()
	fmt.Println()

	fmt.Println("Example 3")
	ex3()
	fmt.Println()

	fmt.Println("Example 4")
	ex4()
}
