package main

import (
	"belajar-gin/router"
	"database/sql"
	"fmt"
	"log"
	_ "net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	db, err := connectMysql()
	if err != nil {
		log.Println(err)
	}

	defer db.Close()
	r := router.SetupRouter()

	r.Run(":8001")
}

func connectMysql() (*sql.DB, error) {
	err := godotenv.Load()
	config := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"),
		os.Getenv("DBNET"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	)
	db, err := sql.Open("mysql", config)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(time.Minute * 5)
	db.SetConnMaxLifetime(time.Minute * 60)

	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
