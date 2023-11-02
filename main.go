package main

import (
	"advisor/database"
	"advisor/routes"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//go:embed templates/*
var templateFS embed.FS

func main() {
	// load .env to get SERVERPORT
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error in loading .env")
	}

	serverPort := os.Getenv("SERVERPORT")
	strServerPort := fmt.Sprintf(":%s", serverPort)

	db := database.NewConnection()
	routes := routes.GetRoutes(db, templateFS)

	log.Printf("application is running on port %s\n", strServerPort)
	http.ListenAndServe(strServerPort, routes)
}
