package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type Record struct {
	Date                                     time.Time
	Open, High, Low, Close, Volume, AdjClose string
}

type Table struct {
	Fields  []string
	Records []Record
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funcs).ParseFiles("index.gohtml"))
}

var funcs = template.FuncMap{
	"fmtDate": fmtDate,
}

func fmtDate(date time.Time) string {
	return date.Format("Jan 2 2006")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func convertAndCheck(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	// Will exit program if error occurs
	checkErr(err, "problem parsing float:")
	return f
}

func main() {
	file, err := os.Open("table.csv")
	checkErr(err, "couldn't open file:")
	defer file.Close()

	reader := csv.NewReader(file)
	table := Table{
		Records: make([]Record, 0), // Fields added below
	}

	// Get first line for Fields
	table.Fields, err = reader.Read()
	checkErr(err, "problem getting first line:")

	// Get rest of records
	for {
		xs, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Check err for other errors
		checkErr(err, "problem reading a record:")

		// Make a Record
		date, err := time.Parse("2006-01-02", xs[0])
		checkErr(err, "problem parsing date:")
		record := Record{
			Date:     date,
			High:     xs[1],
			Open:     xs[2],
			Low:      xs[3],
			Close:    xs[4],
			Volume:   xs[5],
			AdjClose: xs[6],
		}

		// Add to table
		table.Records = append(table.Records, record)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "index.gohtml", table)
	checkErr(err, "problem with tpl.Execute:")
}
