package format

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

// CSV offers support for CSV format.
type CSV struct {
	reader *csv.Reader
}

// Initialize initializes the reader with a new csv.Reader
func (c *CSV) Initialize(file *os.File) error {
	c.reader = csv.NewReader(file)
	c.reader.FieldsPerRecord = -1
	c.reader.LazyQuotes = true
	return nil
}

// ReadHeader reads the header of the file
func (c *CSV) ReadHeader() ([]string, error) {
	headers, err := c.reader.Read()
	if err == io.EOF {
		return headers, errors.New("file must contain header row and at least one data row")
	}
	return headers, nil
}

// ReadRow reads a single CSV row and returns the data as a map of key/value pairs where key is column name and value is row value.
func (c *CSV) ReadRow(headers []string) (map[string]string, error) {
	record, err := c.reader.Read()
	if err != nil {
		return nil, err
	}

	row := make(map[string]string, len(headers))

	for i, header := range headers {
		row[header] = record[i]
	}
	return row, nil
}
