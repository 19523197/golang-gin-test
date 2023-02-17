package main

import (
	"belajar-gin/router"
	"database/sql"
	"fmt"
	"log"
	_ "net/http"
	"os"

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
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}
