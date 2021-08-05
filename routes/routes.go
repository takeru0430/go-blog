package routes

import "net/http"
import "local.pkg/ctrl"
import "os"

// Set ルーティングをセット
func Set() {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")

    /**
     * gb 固有のルート
     */
    http.HandleFunc(p + "/gb-admin", ctrl.Admin) // ダッシュボード
    // 記事関連
    http.HandleFunc(p + "/gb-admin/post/create", ctrl.PostCreateForm)
    http.HandleFunc(p + "/gb-admin/post/create/check", ctrl.PostCreateCheck)
    http.HandleFunc(p + "/gb-admin/post/create/submit", ctrl.PostCreateSubmit)

    http.Handle(p + "/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))

    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "public/favicon.ico")
    })

    // トップページ
    http.HandleFunc(p + "/", ctrl.Home)

    // エラーハンドリングをしていないため強制的にtrue
    //return true
}
