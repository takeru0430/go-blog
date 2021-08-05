package catcherr

import (
	"log"
	"runtime"
)

// CatchAlertError Fatalレベルでエラーを確認 処理を止める
func CatchAlertError(e error) {
	_, file, line, _ := runtime.Caller(1)
	if e != nil {
		log.Fatalf("alert! file: %s: %d, \nmessage: %v", file, line, e)
	}
}

// CatchWarnError Warnレベルでエラーを確認 処理を止めない
func CatchWarnError(e error) {
	_, file, line, _ := runtime.Caller(1)
	if e != nil {
		log.Printf("warning! file: %s:%d, \nmessage: %v", file, line, e)
	}
}

