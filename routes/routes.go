package routes

import "net/http"
import "../ctrl"

/*
 * ルーティングをセット
 */
func Set() bool {
    // トップページ
    http.HandleFunc("/", ctrl.Home)

    // 記事のCURD
    http.HandleFunc("/post/", ctrl.Post)
    http.HandleFunc("/post/create", ctrl.Post)
    http.HandleFunc("/post/edit", ctrl.Post)

    // カテゴリーページ
    http.HandleFunc("/category", ctrl.Category)

    // エラーハンドリングをしていないため強制的にtrue
    return true
}
