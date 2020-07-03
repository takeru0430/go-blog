package ctrl

import "log"
import "net/http"
import . "go-blog/db"

func Home(w http.ResponseWriter, r *http.Request) {
    // パスが / だった場合
    if r.URL.Path == "/" {
        // DB
        row, err := Db.Query("SELECT name FROM `dev_go-blog`.post WHERE id = 1")
        if err != nil {
            log.Fatal("Custom Error : 接続失敗？ %v", err)
        }
        row.Next()
        var name string
        row.Scan(&name)

        // indexビューを取得
        page := loadView("index")

        // indexビューにデータを埋め込み
        page.Execute(w,
            MetaData{
                Title:       "golog｜golangブログフレームワーク",
                Description: name,
                CurrentPath: ".",
            })
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusNotFound)
    }
}
