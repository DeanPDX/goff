package format

import "os"

// Writable is for file types that can be written to.
type Writable interface {
	// WriteHeader should be called at the beginning of the file write process.
	WriteHeader(file *os.File, columns []string) error
	// WriteRecord should be called for every record in the source data.
	WriteRecord(file *os.File, row map[string]string, index int) error
	// WriteFooter should be called after every record has been written with WriteRecord.
	WriteFooter(file *os.File) error
}
