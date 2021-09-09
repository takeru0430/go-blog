package main

import (
    "net/http"
    "os"
)

//import . "go-blog/db"

// home 静的ファイルのコントローラ
func home(w http.ResponseWriter, r *http.Request) {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")

    if r.URL.Path ==  p + "/" {
        // DB
        /*
        row, err := Db.Query("SELECT name FROM `dev_go-blog`.post WHERE id = 1")
        if err != nil {
            log.Fatal("Custom Error : 接続失敗？ %v", err)
        }
        var name string
        row.Scan(&name)
        */

        // indexビューを取得
        page := loadView("index")

        // indexビューにデータを埋め込み
        page.Execute(w,
            MetaData{
                Title:       "go-blog｜ブログフレームワーク",
                Description: "GO言語でWPのようなタクソノミー構造と、Laravelのようなディレクト構造を持ったPHPerライクなブログフレームワークです。",
                CurrentPath: ".",
            })
    } else {
        // ここまで来たリクエストは、存在しないパス
        http.NotFound(w, r)
    }
}
