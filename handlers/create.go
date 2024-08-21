package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

func CreateUserForm(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// フォームデータからID、名前、年齢を取得
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil || id <= 0 {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil || age < 0 {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		// ユーザー情報をデータベースに挿入
		_, err = db.Exec("INSERT INTO users (id, column1, column2) VALUES (?, ?, ?)", id, name, age)
		if err != nil {
			http.Error(w, "Error inserting data", http.StatusInternalServerError)
			return
		}

		// 作成後、ユーザー一覧ページにリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
