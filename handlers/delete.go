// package handlers

// import (
// 	"database/sql"
// 	"html/template"
// 	"net/http"
// )

// func DeleteUserForm(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		id := r.URL.Query().Get("id")

// 		row := db.QueryRow("SELECT Name, Age FROM User_Info WHERE User_ID = ?", id)

// 		var name string
// 		var age int
// 		err := row.Scan(&name, &age)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		data := struct {
// 			ID   string
// 			Name string
// 			Age  int
// 		}{
// 			ID:   id,
// 			Name: name,
// 			Age:  age,
// 		}

// 		tmpl.ExecuteTemplate(w, "delete.html", data)
// 	}
// }

// func DeleteUser(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodPost {
// 			id := r.FormValue("id")

// 			_, err := db.Exec("DELETE FROM User_Info WHERE User_ID = ?", id)
// 			if err != nil {
// 				http.Error(w, err.Error(), http.StatusInternalServerError)
// 				return
// 			}

//				http.Redirect(w, r, "/", http.StatusSeeOther)
//			}
//		}
//	}
//
// ------------------------------------------------------------------------------
package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

func DeleteUserForm(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
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

		tmpl.ExecuteTemplate(w, "delete.html", data)
	}
}

// ユーザーの削除処理
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			id := r.FormValue("id")

			_, err := db.Exec("DELETE FROM User_Info WHERE User_ID = ?", id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"result": "成功しました"})
		}
	}
}
