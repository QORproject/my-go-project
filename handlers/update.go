package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"strconv"
)

func UpdateUserForm(db *sql.DB, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// URLパラメータからIDを取得
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		// 指定されたIDのユーザー情報を取得
		var user struct {
			ID      int
			Column1 string
			Column2 int
		}
		err = db.QueryRow("SELECT id, column1, column2 FROM users WHERE id = ?", id).Scan(&user.ID, &user.Column1, &user.Column2)
		if err != nil {
			http.Error(w, "Error fetching user", http.StatusInternalServerError)
			return
		}

		// 更新フォームにユーザー情報を埋め込んで表示
		err = templates.ExecuteTemplate(w, "update.html", user)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// フォームデータからID、名前、年齢を取得
		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		// ユーザー情報を更新
		_, err = db.Exec("UPDATE users SET column1 = ?, column2 = ? WHERE id = ?", name, age, id)
		if err != nil {
			http.Error(w, "Error updating data", http.StatusInternalServerError)
			return
		}

		// 更新後、ユーザー一覧ページにリダイレクト
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
