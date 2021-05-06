// Package json contains implementation of format interfaces for JSON document type.
package json

import (
	"encoding/json"
	"os"
)

// Writer offers support for JSON format.
type Writer struct{}

// WriteHeader in the case of JSON writes opening array rune: "[".
func (j *Writer) WriteHeader(file *os.File, columns []string) error {
	// Opening of our json array.
	_, err := file.WriteString("[\n\t")
	return err
}

// WriteRecord encodes and writes a record to file.
func (j *Writer) WriteRecord(file *os.File, row map[string]string, index int) error {
	// Encode our object in JSON with indentation
	encoded, err := json.MarshalIndent(row, "\t", "\t")
	if err != nil {
		return err
	}

	// If not first record, add comma, newline and tab to keep our format looking correct in
	// array of JSON data.
	if index > 0 {
		_, err := file.WriteString(",\n\t")
		if err != nil {
			return err
		}
	}

	_, err = file.Write(encoded)
	return err
}

// WriteFooter writes the JSON closing array rune to the footer.
func (j *Writer) WriteFooter(file *os.File) error {
	// Closing of our json array.
	_, err := file.WriteString("\n]")
	return err
}
