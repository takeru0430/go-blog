package ctrl

import "log"
import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
    log.Printf(r.URL.Path)

    // パスが / だった場合
    if r.URL.Path == "/" {
        // indexビューを取得
        page := loadView("index")

        // indexビューにデータを埋め込み
        page.Execute(w,
            MetaData{
                Title:       "golog｜golangブログフレームワーク",
                Description: "説明文",
                CurrentPath: ".",
            })
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusNotFound)
    }
}
