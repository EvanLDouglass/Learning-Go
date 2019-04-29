// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on channels
// Author: Evan Douglass
// Created: April 28 2019

package main

import (
	"fmt"
)

// Given starting code, made this function work
// (caused deadlock initially)
func ex1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

// Given starting code, made this function work
func ex2() {
	// was receive only
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
	fmt.Println()

	// was give only
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)

	fmt.Printf("------\n")
	fmt.Printf("cr\t%T\n", cr)
}

func ex3() {
	c := gen()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func ex4() {
	q := make(chan int)
	c := genEx4(q)

	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func genEx4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}

func ex5() {
	c := make(chan int)
	go func() {
		c <- 1
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

func ex6() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go putOn(c)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(<-c)
	}
}

func putOn(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
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
}
