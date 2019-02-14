package main

import (
	"testing"
)

func TestParseLog(t *testing.T) {

	sampleRow := Row{ID: 0, Message: "hogehoge", Fdesc: "stdout", Time: "2018-09-14"}

	row, _ := parseLog(`{"log": "hogehoge", "stream": "stdout", "time": "2018-09-14"}`)
	if row != sampleRow {
		t.Fatalf("result of parseLog is worry\n%v\n%v", row, sampleRow)
	}

	row, _ = parseLog(`{"log": "hogehoge", "stream": "stdout", "time": "2018-09-14"}`)
	if row.ID != 1 {
		t.Fatalf("id is not incremented %d", row.ID)
	}
}
