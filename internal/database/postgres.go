package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

// ConnectPostgres initializes the PostgreSQL connection
func ConnectPostgres() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("🚨 Error connecting to DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("🚨 Cannot reach DB: %v", err)
	}

	log.Println("✅ Connected to PostgreSQL database successfully")
}

// GetDB returns the database connection
func GetDB() *sql.DB {
	return DB
}
