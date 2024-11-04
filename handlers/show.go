package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

// User構造体
type User struct {
	UserID   int
	Name     string
	Gender   string
	Age      int
	Email    string
	CreateAt string
	UpdateAt string
}

// Book構造体
type Book struct {
	BookID      int
	Title       string
	Author      string
	ReleaseDate string
	Synopsis    string
	CreateAt    string
	UpdateAt    string
}

// ユーザー一覧表示
func ShowUsers(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT User_ID, Name, Gender, Age, Email, CreateAt, UpdateAt FROM User_Info")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []User
		for rows.Next() {
			var user User
			err := rows.Scan(&user.UserID, &user.Name, &user.Gender, &user.Age, &user.Email, &user.CreateAt, &user.UpdateAt)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error scanning users", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		err = tmpl.ExecuteTemplate(w, "users.html", users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// 本一覧表示
func ShowBooks(db *sql.DB, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT Book_ID, Title, Author, ReleaseDate, Synopsis, CreateAt, UpdateAt FROM Book_Info")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var books []Book
		for rows.Next() {
			var book Book
			err := rows.Scan(&book.BookID, &book.Title, &book.Author, &book.ReleaseDate, &book.Synopsis, &book.CreateAt, &book.UpdateAt)
			if err != nil {
				log.Println(err)
				http.Error(w, "Error scanning books", http.StatusInternalServerError)
				return
			}
			books = append(books, book)
		}

		err = tmpl.ExecuteTemplate(w, "books.html", books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
