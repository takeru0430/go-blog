package main

import (
	"log"
	"runtime"
)

// catchAlertError Fatalレベルでエラーを確認 処理を止める
func catchAlertError(e error) {
	_, file, line, _ := runtime.Caller(1)
	if e != nil {
		log.Fatalf("alert! file: %s: %d, \nmessage: %v", file, line, e)
	}
}

// catchWarnError Warnレベルでエラーを確認 処理を止めない
func catchWarnError(e error) {
	_, file, line, _ := runtime.Caller(1)
	if e != nil {
		log.Printf("warning! file: %s:%d, \nmessage: %v", file, line, e)
	}
}

