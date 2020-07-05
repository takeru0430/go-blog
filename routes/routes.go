package routes

import "net/http"
import "go-blog/ctrl"

/*
 * ルーティングをセット
 */
func Set() bool {

    // 記事関連
    http.HandleFunc("/post/create/form", ctrl.PostCreateForm)
    http.HandleFunc("/post/create/check", ctrl.PostCreateCheck)
    http.HandleFunc("/post/create/submit", ctrl.PostCreateSubmit)

    /**
     * admin関連
     */
    http.HandleFunc("/admin", ctrl.Admin) /* ダッシュボード */

    // トップページ
    http.HandleFunc("/", ctrl.Home)
    // エラーハンドリングをしていないため強制的にtrue
    return true
}
