package main

import (
	"net/http"
    "os"
)

// setRouter ルーティングをセット
func setRouter() {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")

    /**
     * gb 固有のルート
     */
    http.HandleFunc(p + "/gb-admin", admin) // ダッシュボード
    // 記事関連
    http.HandleFunc(p + "/gb-admin/post/create", postCreateForm)
    http.HandleFunc(p + "/gb-admin/post/create/check", postCreateCheck)
    http.HandleFunc(p + "/gb-admin/post/create/submit", postCreateSubmit)

    http.Handle(p + "/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))

    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "public/favicon.ico")
    })

    // トップページ
    http.HandleFunc(p + "/", home)

    // エラーハンドリングをしていないため強制的にtrue
    //return true
}
