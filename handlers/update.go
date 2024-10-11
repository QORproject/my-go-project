package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"
)

func UpdateUserForm(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		row := db.QueryRow("SELECT Name, Age FROM User_Info WHERE User_ID = ?", id)

		var name string
		var age int
		err := row.Scan(&name, &age)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := struct {
			ID   string
			Name string
			Age  int
		}{
			ID:   id,
			Name: name,
			Age:  age,
		}

		tmpl.ExecuteTemplate(w, "update.html", data)
	}
}

func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			r.ParseForm()
			id := r.FormValue("id")
			name := r.FormValue("name")
			age := r.FormValue("age")

			_, err := db.Exec("UPDATE User_Info SET Name = ?, Age = ?, UpdateAt = ? WHERE User_ID = ?", name, age, time.Now(), id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}
