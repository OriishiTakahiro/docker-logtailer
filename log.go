package main

import (
	"encoding/json"
)

var (
	currentID = 0
)

// Row represents docker log format
// ex.{"log":"# ls\r\n","stream":"stdout","time":"2019-01-11T08:01:53.36030375Z"}
type Row struct {
	ID      int
	Message string `json:"log"`    // log message
	Fdesc   string `json:"stream"` // file descriptor
	Time    string `json:"time"`   // timestamp
}

// NewRow instantiate Row
func NewRow() Row {
	row := Row{
		ID: currentID,
	}
	currentID++
	return row
}

func parseLog(log string) (Row, error) {
	row := NewRow()
	err := json.Unmarshal([]byte(log), &row)
	if err != nil {
		return row, err
	}
	return row, nil
}
