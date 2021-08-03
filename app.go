package main

import (
	env "github.com/joho/godotenv"
	"local.pkg/routes"
	log "log"
	"net/http"
)

func main() {
    env.Load()

    // ルーティングの初期化
    routes.Set()

    // listen
    err := http.ListenAndServe(":3000", nil)
	if err != nil {
		// log.Exit("ListenAndServe: ", err.String())
        log.Fatal(err)
	}
}
