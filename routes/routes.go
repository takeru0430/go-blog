package routes

import "net/http"
import "go-blog/ctrl"
import "os"
//import "flag"
//import "regexp"
//import "github.com/google/logger"

//const logPath = "./log/go-blog.log"

//var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

/* * ルーティングをセット */
// func Set() bool {
func Set() {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")
    //p = "/" + p


    // 記事関連
    http.HandleFunc(p + "/post/create/form", ctrl.PostCreateForm)
    http.HandleFunc(p + "/post/create/check", ctrl.PostCreateCheck)
    http.HandleFunc(p + "/post/create/submit", ctrl.PostCreateSubmit)

    /**
     * admin関連
     */
    http.HandleFunc(p + "/admin", ctrl.Admin) /* ダッシュボード */

    // トップページ
    http.HandleFunc(p + "/", ctrl.Home)

    // エラーハンドリングをしていないため強制的にtrue
    //return true
}
