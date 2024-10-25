package handlers

import (
	"html/template"
	"net/http"
)

// ホーム画面のハンドラ
func Home(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	}
}
