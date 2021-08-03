package routes

import "net/http"
import "local.pkg/ctrl"
import "os"

// Set ルーティングをセット
func Set() {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")


    // 記事関連
    http.HandleFunc(p + "/post/create/form", ctrl.PostCreateForm)
    http.HandleFunc(p + "/post/create/check", ctrl.PostCreateCheck)
    http.HandleFunc(p + "/post/create/submit", ctrl.PostCreateSubmit)

    /**
     * admin関連
     */
    http.HandleFunc(p + "/admin", ctrl.Admin) /* ダッシュボード */

    http.Handle(p + "/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))

    http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "public/favicon.ico")
    })

    // トップページ
    http.HandleFunc(p + "/", ctrl.Home)

    // エラーハンドリングをしていないため強制的にtrue
    //return true
}
