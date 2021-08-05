package db

import "database/sql"
import "log"
import "os"
import _ "github.com/go-sql-driver/mysql"
import env "github.com/joho/godotenv"

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

// Db 外部からこの接続を利用してSQLを流せるようにする
var Db *sql.DB = connect()

