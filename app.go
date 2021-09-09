package main

import (
	"log"
	"net/http"
)

func main() {
    // ルーティングの初期化
    setRouter()

    // listen
    err := http.ListenAndServe(":3000", nil)
	if err != nil {
		// log.Exit("ListenAndServe: ", err.String())
        log.Fatal(err)
	}
}
