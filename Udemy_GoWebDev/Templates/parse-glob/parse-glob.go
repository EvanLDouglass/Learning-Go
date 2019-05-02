// Evan Douglass
// Created May 01 2019
// Using parse glob

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Do on load
func init() {
	// Must checks the err for me
	tpl = template.Must(template.ParseGlob("templates/*.txt"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "three.txt", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
