package main

import (
	"advisor/database"
	"advisor/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env to get SERVERPORT
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error in loading .env")
	}

	serverPort := os.Getenv("SERVERPORT")
	strServerPort := fmt.Sprintf(":%s", serverPort)

	db := database.NewConnection()
	routes := routes.GetRoutes(db)

	log.Printf("application is running on port %s\n", strServerPort)
	http.ListenAndServe(strServerPort, routes)
}
