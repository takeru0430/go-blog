package main

import "net/http/cgi"
import "go-blog/routes"
import env "github.com/joho/godotenv"

func main() {
    env.Load()

    // ルーティングの初期化
    routes.Set()

    // listen
    cgi.Serve(nil)
}
