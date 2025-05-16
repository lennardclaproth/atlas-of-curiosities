package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Event struct {
	ID            string
    StreamID      string
    Type          string
    Timestamp     int64
    CorrelationID string
    CausationID   string
    Data          []byte
    Metadata      []byte
}

func main() {
	d1 := []byte("hello\ngo\n")
	path, err := os.Getwd()
	dir := "tmp"

	if err := os.Mkdir(filepath.Join(path, dir), 0755); os.IsExist(err) {
		fmt.Println("The directory named", dir, "exists")
	 } else {
		fmt.Println("The directory named", dir, "does not exist")
	 }

	err = os.WriteFile(filepath.Join(path, "tmp", "1.cdb"), d1, 0644)

	if err != nil {
		// Do smth with err.
	}
}

func write_event_bytes() {
	
}
