// Simple exercises to get familiar with the Go programming language.
// Exercises are from the Udemy course Learn How To Code Google's Go.
// This file focuses on highlighted packages from the course videos.
// Author: Evan Douglass
// Created: April 27 2019

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// User is a simple struct for the JSON exercise
type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// Marshaling JSON
func ex1() {
	// Users
	u1 := User{"James", "Bond", 32, nil}
	u2 := User{"Miss", "Moneypenny", 27, nil}
	u3 := User{"M", ".", 54, nil}

	users := []User{u1, u2, u3}
	fmt.Println(users)

	js, err := json.Marshal(users)

	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(string(js))
	}
}

// Un-marshaling JSON
func ex2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var users []User
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%+v\n", users)
	}
}

// More JSON, with encoder
func ex3() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	enc := json.NewEncoder(os.Stdout) // returns pointer
	err := enc.Encode(users)
	if err != nil {
		fmt.Println("error:", err)
	}
}

// Using sort
func ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort xi
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort xs
	sort.Strings(xs)
	fmt.Println(xs)
}

// Allow sorting by User.Age
type byAge []User

// Methods for byAge
func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

// Allow sorting by User.Last
type byLast []User

// Methods for byLast
func (l byLast) Len() int {
	return len(l)
}

func (l byLast) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l byLast) Less(i, j int) bool {
	return l[i].Last < l[j].Last
}

// Pretty print a slice of users
func printUsers(u []User) {
	for _, v := range u {
		fmt.Println(v.First, v.Last, v.Age)
		for _, v := range v.Sayings {
			fmt.Printf("\t%s\n", v)
		}
	}
}

// Sorting structs
func ex5() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	printUsers(users)
	fmt.Println()

	// Sort each user's sayings
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	printUsers(users)
	fmt.Println()

	sort.Sort(byAge(users))
	printUsers(users)
	fmt.Println()

	sort.Sort(byLast(users))
	printUsers(users)
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
}
