// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on concurrency in Go
// Author: Evan Douglass
// Created: April 27 2019

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Using go routines with a WaitGroup
func ex1() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("Hello from 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Hello from 2")
	}()

	wg.Wait()
}

// For Ex. 2
type person struct {
	first    string
	last     string
	greeting string
	age      int
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println(p.greeting)
}

func saySomething(h human) {
	h.speak()
}

func ex2() {
	me := person{
		first:    "Evan",
		last:     "Douglass",
		greeting: "Howdy!",
		age:      28,
	}

	// Can pass pointer into saySomething
	saySomething(&me)

	// Can't pass by actual value because of speak method
	// This causes an error
	//saySomething(me)
}

// Counts to 1000 with a race condition
func ex3() {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// atomically add value
			counter++
			// decrement waitgroup
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

// Ex4 counts to 1000 using a mutex
func Ex4() {
	var counter int64 // int64 used to compare time with next function
	var mut sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// lock and add a value
			mut.Lock()
			counter++
			mut.Unlock()
			// decrement waitgroup
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

// Ex5 counts to 1000 using atomic package
func Ex5() {
	var counter int64
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			// atomically add value
			atomic.AddInt64(&counter, 1)
			// decrement waitgroup
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}

// Using package runtime constants
func ex6() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumCPU())
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
	Ex4()
	fmt.Println()

	fmt.Println("Exercise 5")
	Ex5()
	fmt.Println()

	fmt.Println("Exercise 6")
	ex6()
	fmt.Println()
}
