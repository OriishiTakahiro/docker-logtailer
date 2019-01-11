package main

import (
	"strings"
)

func init() {
	registerMsgPatterns(".+: No such file or directory", fileNotFound)
}

func fileNotFound(row Row, matched []string) string {

	msg := row.Message +
		`
----- エラーの種類 -----
ファイル及びディレクトリがありません

----- 対策案 -----
対象のファイル，ディレクトリがあるか確認してください．
対象のファイル，ディレクトリが存在する

----- 対象ファイル/ディレクトリ ----
	`

	for _, file := range matched {
		msg = msg + "- " + strings.Replace(file, ": No such file or directory", "", -1) + "\n"
	}

	return msg
}
