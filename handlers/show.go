package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func ShowUsers(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT User_ID, Name, Age FROM User_Info")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		tmpl.ExecuteTemplate(w, "users.html", struct{ Users []User }{Users: users})
	}
}
