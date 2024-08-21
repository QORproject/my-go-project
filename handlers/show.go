package handlers

import (
	"database/sql"
	"html/template"
	"net/http"
)

func ShowUsers(db *sql.DB, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, column1, column2 FROM users")
		if err != nil {
			http.Error(w, "Error querying database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []struct {
			ID      int
			Column1 string
			Column2 int
		}

		for rows.Next() {
			var user struct {
				ID      int
				Column1 string
				Column2 int
			}
			err := rows.Scan(&user.ID, &user.Column1, &user.Column2)
			if err != nil {
				http.Error(w, "Error scanning data", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		err = templates.ExecuteTemplate(w, "users.html", map[string]interface{}{
			"Users": users,
		})
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
	}
}
