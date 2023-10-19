package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func NewConnection() *sql.DB {
	// load .env file to get db connection
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error in loading .env file")
	}

	host := os.Getenv("POSTGRESHOST")
	port := os.Getenv("POSTGRESPORT")
	user := os.Getenv("POSTGRESUSER")
	password := os.Getenv("POSTGRESPASSWORD")
	database := os.Getenv("POSTGRESDATABASE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("error accessing database: %v\n", err)
	}

	// ping before returning connection
	if err := connection.Ping(); err != nil {
		log.Fatalln("database connection error. Please check PostgreSQL Service")
		log.Fatalf("cannot connect to database: %v\n", err)
	}

	log.Printf("database connected on port %s\n", port)
	return connection
}
