package main

import (
	"local.pkg/routes"
	log "log"
	"net/http"
)

func main() {
    // ルーティングの初期化
    routes.Set()

    // listen
    err := http.ListenAndServe(":3000", nil)
	if err != nil {
		// log.Exit("ListenAndServe: ", err.String())
        log.Fatal(err)
	}
}
