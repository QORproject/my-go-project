package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// MySQLへの接続文字列 (DSN)
	dsn := "root:Kanta0930@tcp(db:3306)/mydb"

	// MySQLへの接続を試みる (最大10回リトライ)
	var db *sql.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Failed to connect to MySQL: %v\n", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Pingを送って接続確認
		err = db.Ping()
		if err == nil {
			fmt.Println("Connected to MySQL!")
			break
		}

		log.Printf("MySQL connection failed: %v. Retrying...\n", err)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	defer db.Close()

	// データの挿入
	_, err = db.Exec("INSERT INTO users (column1, column2) VALUES (?, ?)", "John Doe", 30)
	if err != nil {
		log.Fatalf("Error inserting data: %v", err)
	}
	fmt.Println("Data inserted successfully!")

	// データの取得
	rows, err := db.Query("SELECT id, column1, column2 FROM users")
	if err != nil {
		log.Fatalf("Error querying data: %v", err)
	}
	defer rows.Close()

	fmt.Println("Data from 'users' table:")
	for rows.Next() {
		var id int
		var column1 string
		var column2 int
		err := rows.Scan(&id, &column1, &column2)
		if err != nil {
			log.Fatalf("Error scanning data: %v", err)
		}
		fmt.Printf("ID: %d, Column1: %s, Column2: %d\n", id, column1, column2)
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		log.Fatalf("Error with rows: %v", err)
	}
}
