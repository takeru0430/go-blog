package ctrl

import "html/template"
import "log"
import "net/http"

// 各ページで共通の構造体
type MetaData struct {
    Title string
    Description string
    CurrentPath string
    FlashText string
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
        "../view/_footer.html",
    )

    if err != nil {
        log.Fatalf("template Error: %v", err)
    }

    return page
}

