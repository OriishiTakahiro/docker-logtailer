package main

import (
	"regexp"
)

var (
	msgPatterns = []MsgPattern{}
)

// MsgPattern represents pair of Regexp pattern and message generator
type MsgPattern struct {
	Pattern      *regexp.Regexp
	MsgGenerator func(Row, []string) string
}

func generateMsg(row Row) (string, bool) {
	for _, v := range msgPatterns {
		matched := v.Pattern.FindAllString(row.Message, -1)
		if matched != nil {
			return v.MsgGenerator(row, matched), true
		}
	}
	return "", false
}

func registerMsgPatterns(pattern string, generator func(Row, []string) string) error {

	compiled, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	msgPatterns = append(msgPatterns, MsgPattern{Pattern: compiled, MsgGenerator: generator})
	return nil
}
