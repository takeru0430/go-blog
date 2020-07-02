package main

import "log"
import "net/http/cgi"
import "./routes"

func main() {
    //http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("public/"))))

    // ルーティングのセット
    if routes.Set() != true {
        log.Printf("エラー：ルーティングに問題があります。")
    }

    // listen
    cgi.Serve(nil)
}
