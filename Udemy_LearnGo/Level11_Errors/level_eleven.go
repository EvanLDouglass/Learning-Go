// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on errors
// Author: Evan Douglass
// Created: April 29 2019

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"time"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func ex1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Updated to check for error using the Go idiom of a second return value
	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("error in json.Marshal: %v", err)
	}
	fmt.Println(string(bs))
}

func ex2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// Updated to check for error using the Go idiom of a second return value
	// Done with fmt.Errorf in this function
	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		err = fmt.Errorf("error in toJSON: %v", err)
	}
	return bs, err
}

// For Ex. 3
type customErr struct {
	t   time.Time
	msg string
}

func (e customErr) Error() string {
	return fmt.Sprintf("%v: %v", e.t, e.msg)
}

func ex3() {
	var myErr customErr
	myErr.msg = "The sky is falling!"
	myErr.t = time.Now()
	causeErr(myErr)
}

func causeErr(e error) {
	fmt.Println(e)
}

// For Ex. 4
type sqrtError struct {
	lat  float64
	long float64
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func ex4() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// write your error code here
		err := sqrtError{
			lat:  47.614928299999995,
			long: -122.3401989, // Seattle
			err:  errors.New("can't square root a negative"),
		}
		return 0, err
	}
	return math.Sqrt(f), nil
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
}
