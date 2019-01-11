package main

import (
	"github.com/hpcloud/tail"
)

// ファイルを追跡しながら読み出す
func tailFile(msgChan chan string, errChan chan error, filename string) {
	// ファイルを追跡
	t, err := tail.TailFile(filename, tail.Config{Follow: true})
	// エラーチェック
	if err != nil {
		errChan <- err
	}
	for line := range t.Lines {
		msgChan <- line.Text
	}
}
