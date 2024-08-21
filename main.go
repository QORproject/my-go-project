package main

import (
	"database/sql"
	"html/template"
	"log"
	"my-go-project/handlers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var templates *template.Template

func init() {
	// テンプレートを一括でパースする
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// MySQLへの接続文字列 (DSN)
	dsn := "root:Kanta0930@tcp(db:3306)/mydb"

	// MySQLへの接続を試みる
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}
	defer db.Close()

	// MySQL接続確認
	err = db.Ping()
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	log.Println("Connected to MySQL!")

	// エンドポイントの設定
	http.HandleFunc("/", handlers.ShowUsers(db, templates))
	http.HandleFunc("/create", handlers.CreateUserForm(templates))
	http.HandleFunc("/create/submit", handlers.CreateUser(db))
	http.HandleFunc("/update", handlers.UpdateUserForm(db, templates))
	http.HandleFunc("/update/submit", handlers.UpdateUser(db))
	http.HandleFunc("/delete", handlers.DeleteUserForm(db, templates))
	http.HandleFunc("/delete/submit", handlers.DeleteUser(db))

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
