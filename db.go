package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    env "github.com/joho/godotenv"
    "log"
    "os"
)

func connect() *sql.DB {
    err := env.Load(".env")
    if err != nil {
        log.Fatal(err)
    }

    // 環境変数の読み込みとセット
    u := os.Getenv("DB_USER")
    p := os.Getenv("DB_PASS")
    h := os.Getenv("DB_HOST")
    n := os.Getenv("DB_NAME")

    // DBに接続
    db, err := sql.Open("mysql", u + ":" + p + "@tcp(" + h + ")/" + n)
    if err != nil {
        // エラー時の処理
        log.Fatalf("err:DB接続エラー %v", err)
    }

    return db
}

// db 外部からこの接続を利用してSQLを流せるようにする
var db *sql.DB = connect()

