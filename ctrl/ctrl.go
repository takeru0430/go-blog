package ctrl

import "html/template"
import "log"
import "net/http"

// 各ページで共通の構造体
type MetaData struct {
    Title string
    Description string
    CurrentPath string
}

func Category(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/category" {
        // indexビューを取得
        page := loadView("category")

        // indexビューにデータを埋め込み
        page.Execute(w,
            MetaData{
                Title:       "golog｜golangブログフレームワーク",
                Description: "説明文",
                CurrentPath: "..",
            })
    }
}

func Post(w http.ResponseWriter, r *http.Request) {
}

/**
 * 内部関数
 */

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    if status == http.StatusNotFound {
        http.NotFound(w, r)
        log.Printf("custom 404\n")
    }
}

func loadView(n string) *template.Template {
    page, err := template.ParseFiles(
        "../view/" + n + ".html",
        "../view/_header.html",
    )

    if err != nil {
        log.Fatalf("template Error: %v", err)
    }

    return page
}

