/*
* Goff is a quick and easy CLI file formatter meant to convert between formats.
 */
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DeanPDX/goff/format/json"

	"github.com/DeanPDX/goff/format/csv"

	"github.com/DeanPDX/goff/format"
)

// Version is a constant for keeping track of app version.
const Version = "0.1.0"

func main() {
	inPath, outPath := parseArgs(os.Args[1:])
	f, err := os.Open(inPath)
	if err != nil {
		log.Fatal("Problem opening", inPath, ":", err.Error())
	}
	defer f.Close()

	var reader format.Readable
	// TODO: Once we support more file types, swap CSV implementation for whatever.
	reader = &csv.Reader{}

	err = reader.Initialize(f)
	if err != nil {
		log.Fatal(err)
	}

	columns, err := reader.ReadHeader()
	if err != nil {
		log.Fatal(err)
	}

	outFile, _ := os.Create(outPath)
	defer outFile.Close()
	var output format.Writable
	// TODO: Once we support more file types, swap JSON implementation for whatever.
	output = &json.Writer{}
	output.WriteHeader(outFile, columns)
	index := 0
	for {
		row, err := reader.ReadRow(columns)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		output.WriteRecord(outFile, row, index)
		index++
	}
	output.WriteFooter(outFile)
}

func parseArgs(args []string) (string, string) {
	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}

	in, out := args[0], ""

	if len(args) > 1 {
		out = args[1]
	}

	// If out is not supplied, default to appropriate output path.
	if out == "" {
		out = fmt.Sprintf("%v.json", strings.TrimSuffix(in, filepath.Ext(in)))
	}
	return in, out
}

func printHelp() {
	fmt.Printf("v%v\n", Version)
	fmt.Println(`usage:

goff <input file path> <OPTIONAL: output file path>

the second argument (output file path) is optional
and will default to the name of your input file with 
the appropriate extension.`)
}
