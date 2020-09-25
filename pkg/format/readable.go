package format

import "os"

// Readable is an interface for file types that can be read from.
type Readable interface {
	// Initialize should initialize a encoding/<format> reader from the underlying file.
	Initialize(file *os.File) error
	// ReadHeader should return a list of column names.
	ReadHeader() ([]string, error)
	// ReadRow should read a single row and return it as a map[string]string where column name is key and column value is value.
	ReadRow(headers []string) (map[string]string, error)
}
