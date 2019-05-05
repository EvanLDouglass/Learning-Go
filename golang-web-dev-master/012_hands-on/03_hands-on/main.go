package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type CaliHotel struct {
	Name    string
	Address string
	City    string
	Region  string
	Zip     string
}

func main() {
	hotels := []CaliHotel{
		CaliHotel{
			"Hotel California",
			"1234 Cali Lane",
			"LA",
			"Southern",
			"12345",
		},
		CaliHotel{
			"Hotel California North",
			"4321 SanFran St.",
			"San Francisco",
			"Central",
			"99999",
		},
		CaliHotel{
			"La Qinta",
			"0000 Yachty Ave.",
			"Somewhere North",
			"Northern",
			"98765-1234",
		},
	}

	err := tpl.Execute(os.Stdout, hotels)
	if err != nil {
		log.Fatalln(err)
	}
}
