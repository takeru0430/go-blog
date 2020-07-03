package main

import "log"
import "net/http/cgi"
import "go-blog/routes"
import env "github.com/joho/godotenv"

func main() {
    //http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public/"))))
    env.Load()

    // ルーティングのセット
    if routes.Set() != true {
        log.Printf("エラー：ルーティングに問題があります。")
    }

    // listen
    cgi.Serve(nil)
}
