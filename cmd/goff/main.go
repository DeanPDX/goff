package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

var inPath = flag.String("in", "", "The input path of the file you want to parse (Required)")
var outPath = flag.String("out", "", "The output path for the file you want to write (Required)")

func main() {
	flag.Parse()

	if *inPath == "" || *outPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	f, err := os.Open(*inPath)
	if err != nil {
		log.Fatal("Problem opening", *inPath, ":", err.Error())
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	r.LazyQuotes = true

	headers, err := r.Read()
	if err == io.EOF {
		log.Fatal("File must contain header row and at least one data row.")
	}

	outFile, _ := os.OpenFile(*outPath, os.O_CREATE, os.ModePerm)
	defer outFile.Close()
	outFile.WriteString("[\n\t")
	first := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if first == false {
			outFile.WriteString(",\n\t")
		}
		first = false

		columns := make(map[string]string, len(headers))

		for i, header := range headers {
			columns[header] = record[i]
		}

		encoded, err := json.MarshalIndent(columns, "\t", "\t")

		if err != nil {
			log.Fatal(err)
		}
		outFile.Write(encoded)
	}
	outFile.WriteString("\n]")
}
