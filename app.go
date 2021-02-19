package main

import log "log"
import "net/http"
import "local.pkg/routes"
import env "github.com/joho/godotenv"

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
