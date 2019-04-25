// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on functions.
// Author: Evan Douglass
// Created: April 24 2019

package main

import (
	"fmt"
	"math"
)

func foo() int {
	return 5
}

func bar() (int, string) {
	return 5, "five"
}

// Takes a variable number of ints and sums them
func sum(v ...int) int {
	sum := 0
	for _, v := range v {
		sum += v
	}
	return sum
}

// Using defer
func defers() {
	defer fmt.Println("First in func")
	fmt.Println("Second in func")
}

// For Ex. 4
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("My name is %s and I'm %d years old\n",
		p.first, p.age)
}

// For Ex. 5
type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return 2 * math.Pi * c.radius
}

func info(s shape) float64 {
	return s.area()
}

// Returning functions
func returnFunc() func(string) string {
	return func(name string) string {
		return fmt.Sprint("Hello, ", name)
	}
}

// Using a callback
// This function takes a slice of ints and a function
// to perform some cumulative operation on them.
func callback(arr []int, f func([]int) int) int {
	return f(arr)
}

func multiply(arr []int) int {
	total := 1
	for _, v := range arr {
		total *= v
	}
	return total
}

// Using closures
func cumulativeSum() func(float64) float64 {
	total := 0.0
	return func(num float64) float64 {
		total += num
		return total
	}
}

func main() {
	// Ex. 1
	fmt.Println(foo())
	fmt.Println(bar())
	// Ex. 2
	xi := []int{2, 4, 6, 8, 10}
	fmt.Println(sum(xi...))
	// Ex. 3
	defers()
	// Ex. 4
	me := person{
		first: "Evan",
		last:  "Douglass",
		age:   28,
	}
	me.speak()
	// Ex. 5
	sq := square{length: 30, width: 40}
	circ := circle{radius: 20}
	fmt.Println(info(sq))
	fmt.Println(info(circ))
	// Ex. 6 Anonymous function
	fmt.Println(func() string {
		return "Anonymous!"
	}())
	// Ex. 7
	p := fmt.Println
	p("Greetings from a var func")
	// Ex. 8
	hi := returnFunc()
	fmt.Println(hi("Evan"))
	// Ex. 9
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(callback(arr, multiply))
	// Ex. 10
	cSum := cumulativeSum()
	cSum2 := cumulativeSum()
	fmt.Println(cSum(6))
	fmt.Println(cSum(7.7))
	fmt.Println(cSum(0))
	fmt.Println(cSum(-80))
	fmt.Println(cSum2(120))
}
