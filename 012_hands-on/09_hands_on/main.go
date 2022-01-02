package main

import (
    "encoding/csv"    
	"log"
	"os"
	"html/template"
	"time"
	"strconv"	
)


type Record struct {
	Date time.Time
	Open float64
}


func printToWeb(records []Record) {
	
	var tpl *template.Template
	var err error
	//tpl, err = template.Must(template.ParseFiles("hw.gohtml"))
	tpl, err =  template.ParseFiles("hw.gohtml")
	if err != nil {
		log.Fatalln(err)
	}	

	// execute template
	err = tpl.Execute(os.Stdout, records)
	if err != nil {
		log.Fatalln(err)
	}
	
}


func main() {
		// open file
		f, err := os.Open("table.csv")
		if err != nil {
			log.Fatal(err)
		}
	
		// remember to close the file at the end of the program
		defer f.Close()
	
		// read csv values using csv.Reader
		csvReader := csv.NewReader(f)
		rows, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
	
		records := make([]Record, 0, len(rows))

		for i, row := range rows {
			if i == 0 {
				continue
			}
			date, _ := time.Parse("2006-01-02", row[0])
			open, _ := strconv.ParseFloat(row[1], 64)
	
			records = append(records, Record{
				Date: date,
				Open: open,
			})
		}
		
		printToWeb(records)
}

