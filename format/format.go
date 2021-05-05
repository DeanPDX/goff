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

// Writable is for file types that can be written to.
type Writable interface {
	// WriteHeader should be called at the beginning of the file write process.
	WriteHeader(file *os.File, columns []string) error
	// WriteRecord should be called for every record in the source data.
	WriteRecord(file *os.File, row map[string]string, index int) error
	// WriteFooter should be called after every record has been written with WriteRecord.
	WriteFooter(file *os.File) error
}
