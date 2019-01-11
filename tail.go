package main

import (
	"errors"
	"fmt"
	"github.com/hpcloud/tail"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		panic(errors.New("please input target filename"))
	}

	// チャンネルはメッセージを複数のgoroutine間で受け渡し
	msgChan := make(chan string)
	// tailFile()をgroutineとして実行
	go tailFile(msgChan, os.Args[1])

	star := true
	i := 1

	// 無限ループ
	for {
		select {
		// msgChanからメッセージを受け取った際の挙動
		case msg := <-msgChan:
			if msg == "finish" {
				return
			}
			fmt.Printf("[%0d] %s", i, msg)
			i++
		// msgChanからメッセージを受け取らなかった場合の挙動
		default:
			if star {
				fmt.Print("[*]\r")
			} else {
				fmt.Print("[+]\r")
			}
			star = !star
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// ファイルを追跡しながら読み出す
func tailFile(msgChan chan string, filename string) {
	// ファイルを追跡
	t, err := tail.TailFile(filename, tail.Config{Follow: true})
	// エラーチェック
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		msgChan <- line.Text
	}
}
