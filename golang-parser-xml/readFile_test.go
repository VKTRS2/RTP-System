package main

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestReadFile(t *testing.T) {
	for _, path := range fileToRead {
		data := readFile(path)
		decoder := json.NewDecoder(bytes.NewBuffer(data))
		var tr Transfers
		err := decoder.Decode(&tr)
		checkErr(err)
		if len(tr.Transfers) != 3 {
			t.Errorf("Expected to get 3 transactions, but got: %v", len(tr.Transfers))
		}
	}
}
