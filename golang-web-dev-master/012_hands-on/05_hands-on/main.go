package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type Menu struct {
	RName, Meal string
	Items       map[string]float32 // Item to price
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funcs).ParseFiles("index.gohtml"))
}

var funcs = template.FuncMap{
	"asPrice": printPrice,
}

func printPrice(price float32) string {
	return fmt.Sprintf("$%.2f", price)
}

func main() {
	menu := []Menu{
		Menu{
			RName: "McDonalds",
			Meal:  "Breakfast",
			Items: map[string]float32{
				"Pancakes":     5.0,
				"Egg McMuffin": 4.99,
				"Orange Juice": 2.5,
			},
		},
		Menu{
			RName: "McDonalds",
			Meal:  "Lunch - Dinner",
			Items: map[string]float32{
				"Cheese Burger": 3.5,
				"Big Mac":       4.0,
				"Fries":         1.5,
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", menu)
	if err != nil {
		log.Fatalln(err)
	}
}
