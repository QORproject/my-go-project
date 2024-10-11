package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
)

func CreateUserForm(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "create.html", nil)
	}
}

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			name := r.FormValue("name")
			gender := r.FormValue("gender")
			age := r.FormValue("age")
			email := r.FormValue("email")
			password := r.FormValue("password")

			_, err := db.Exec("INSERT INTO User_Info (Name, Gender, Age, Email, Password, CreateAt, UpdateAt) VALUES (?, ?, ?, ?, ?, ?, ?)",
				name, gender, age, email, password, time.Now(), time.Now())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
