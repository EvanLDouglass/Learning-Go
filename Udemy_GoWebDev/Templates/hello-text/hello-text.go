// Evan Douglass
// Created May 01 2019
// A basic approach to templating text

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt") // Used "tpl.gohtml" for first examples
	if err != nil {
		log.Fatalln(err)
	}

	// === Write to stdout below ===
	// err = tpl.Execute(os.Stdout, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// === Create and write to file ===
	// file, err := os.Create("index.html")
	// if err != nil {
	// 	log.Println("error creating file", err)
	// }
	// defer file.Close()

	// err = tpl.Execute(file, nil)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// === Execute Multiple Files ===
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse more files into template
	tpl, err = tpl.ParseFiles("two.txt", "three.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute templates individually by name
	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	// Output first file found in template container
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
