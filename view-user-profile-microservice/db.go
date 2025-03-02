package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database not responding: ", err)
	}

	fmt.Println("Connected to the database successfully!")
}
