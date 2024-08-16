package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
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

	// テーブルが存在するか確認
	_, err = db.Exec("SELECT 1 FROM users LIMIT 1")
	if err != nil {
		if _, ok := err.(*mysql.MySQLError); ok {
			log.Fatalf("Table 'users' doesn't exist: %v", err)
		} else {
			log.Fatalf("Error querying database: %v", err)
		}
	} else {
		fmt.Println("Table 'users' exists!")
	}

}
