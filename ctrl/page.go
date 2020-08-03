package ctrl

//import "log"
import "net/http"
import "os"
//import . "go-blog/db"


/**
 * 固定ページ系のコントローラ
 */

// トップページ
func Home(w http.ResponseWriter, r *http.Request) {
    // get env data
    p := os.Getenv("SUB_DIRECTORY_PATH")

    if r.URL.Path == p + "/" {
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
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}
