// package main

// import (
// 	"database/sql"
// 	"html/template"
// 	"log"
// 	"my-go-project/handlers"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB
// var templates *template.Template

// func init() {
// 	// テンプレートを一括でパースする
// 	templates = template.Must(template.ParseGlob("templates/*.html"))
// }

// func main() {
// 	// MySQLへの接続文字列 (DSN)
// 	dsn := "root:Kanta0930@tcp(db:3306)/mydb"

// 	// MySQLへの接続を試みる
// 	var err error
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MySQL: %v", err)
// 	}
// 	defer db.Close()

// 	// MySQL接続確認
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalf("Could not connect to MySQL: %v", err)
// 	}
// 	log.Println("Connected to MySQL!")

// 	// エンドポイントの設定
// 	http.HandleFunc("/", handlers.ShowUsers(db, templates))            // ユーザー一覧表示
// 	http.HandleFunc("/create", handlers.CreateUserForm(templates))     // ユーザー作成フォーム表示
// 	http.HandleFunc("/create/submit", handlers.CreateUser(db))         // ユーザー作成処理
// 	http.HandleFunc("/update", handlers.UpdateUserForm(db, templates)) // ユーザー更新フォーム表示
// 	http.HandleFunc("/update/submit", handlers.UpdateUser(db))         // ユーザー更新処理
// 	http.HandleFunc("/delete", handlers.DeleteUserForm(db, templates)) // ユーザー削除フォーム表示
// 	http.HandleFunc("/delete/submit", handlers.DeleteUser(db))         // ユーザー削除処理

// 	log.Println("Starting server on :8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

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

	// テンプレートのパース
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	// ホーム画面
	http.HandleFunc("/", handlers.Home(tmpl))

	// ユーザー作成フォームと処理
	http.HandleFunc("/create/user", handlers.CreateUserForm(tmpl))
	http.HandleFunc("/create/user/submit", handlers.CreateUser(db))

	// 本作成フォームと処理
	http.HandleFunc("/create/book", handlers.CreateBookForm(tmpl))
	http.HandleFunc("/create/book/submit", handlers.CreateBook(db))

	//ユーザ一覧、本一覧
	http.HandleFunc("/users", handlers.ShowUsers(db, tmpl))
	http.HandleFunc("/books", handlers.ShowBooks(db, tmpl))

	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
