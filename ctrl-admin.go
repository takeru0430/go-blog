package main

import (
    "net/http"
)
//import "regexp"
//import . "go-blog/db"
import "log"
import "strings"


/**
 * 管理系のコントローラ
 */

// Admin ダッシュボード
func admin(w http.ResponseWriter, r *http.Request) {
    // ビューの取得
    p := loadView("gb/admin/index")

    // データセット
    err := p.Execute(w,
        MetaData{
            Title:       "管理画面｜ブログフレームワーク",
            Description: "",
            CurrentPath: "../..",
        })
    catchWarnError(err)
}

// AdminPostCreate 記事の作成
func AdminPostCreate(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "ga/admin/post/create" {
        if r.Method == http.MethodGet {
            // ビューの取得
            p := loadView("post/create/index")

            // データセット
            p.Execute(w,
                MetaData{
                    Title:       "記事の新規作成｜ブログフレームワーク",
                    Description: "",
                    CurrentPath: "../..",
                })
        } else if r.Method == http.MethodPost {
            log.Print("a")
        }
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}

// AdminPostEdit 記事の編集
func AdminPostEdit(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/admin/post/edit" {
        log.Printf("a")
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}

// AdminPost 記事の一覧 or 特定の記事の詳細
func AdminPost(w http.ResponseWriter, r *http.Request) {
    if strings.Index(r.URL.Path, "/admin/post/1") > 0 {
        log.Printf("a")
    } else {
        // 全体的なエラーを管理
        errorHandler(w, r, http.StatusInternalServerError)
    }
}
