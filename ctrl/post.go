package ctrl

import "net/http"
//import "regexp"
import . "local.pkg/db"
import "log"
//import "strings"

// 記事のデータ構造
type PostData struct {
    PostTitle string
    PostSlug string
    PostBody string
    PostedId string
    MetaData
}



/**
 * 記事系のコントローラ
 */
// 記事の作成
func PostCreateForm(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/post/create/form" {
        // ビューの取得
        p := loadView("post/create/form")

        // データセット
        p.Execute(w,
            MetaData{
                Title:       "記事の新規作成｜ブログフレームワーク",
                Description: "",
                CurrentPath: "../../..",
            })

        return
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}

// 記事の確認
func PostCreateCheck(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/post/create/check" {
        if r.Method == "POST" {
            // ビューの取得
            p := loadView("post/create/check")

            // 構造体の初期化
            PD := PostData{}
            // データのセット
            PD.Title = "記事の確認｜ブログフレームワーク"
            PD.CurrentPath = "../../.."
            PD.PostTitle = r.FormValue("title")
            PD.PostSlug = r.FormValue("slug")
            PD.PostBody = r.FormValue("body")

            // データセット
            p.Execute(w, PD)
        }

        return
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusNotFound)
    }
}

// 記事の保存
func PostCreateSubmit(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/post/create/submit" {
        if r.Method == "POST" {
            // ビューの取得
            p := loadView("post/create/submit")

            // 取得したデータの保存
            result, err := Db.Exec("INSERT INTO post(title, slug, body) VALUES (?, ?, ?)",
                r.FormValue("title"),
                r.FormValue("slug"),
                r.FormValue("body"))

            // 構造体の初期化
            PD := PostData{}
            PD.Title = "記事の公開｜ブログフレームワーク"
            PD.CurrentPath = "../../.."
            // 保存した記事のIDを取得
            if i, err := result.LastInsertId(); err != nil {
                log.Fatal("Error: LastInsertIdメソッドは現在のDBエンジンに対応してません")
                PD.PostedId = "IDは取得不可"
            } else {
                PD.PostedId = string(i)
            }

            // 保存失敗時
            if err != nil {
                log.Fatal("Error: 記事の保存に失敗しました。%v", err)

                PD.FlashText = "記事の公開に失敗しました。"

                p.Execute(w, PD)

            // 保存成功時
            } else {
                // 構造体の初期化
                PD := PostData{}

                PD.FlashText = "記事の公開に成功しました。"

                p.Execute(w, PD)
            }
        }

        return
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}
