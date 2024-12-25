package database

import (
	"anywhere-api/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func DBConnect() *sql.DB {
	// Load the configuration settings
	dbConfig := config.LoadConfig()

	// Prepare the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUser, dbConfig.DBPass, dbConfig.DBName)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
