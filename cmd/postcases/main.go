package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/khaledyd/postcases/internal/app"
	"github.com/khaledyd/postcases/internal/database"
)

func main() {

	// Connect to the PostgreSQL database

	db, err := database.ConnectToPostgres()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer db.Close()
	database.DB = db

	// Create a new Gin instance
	r := gin.Default()

	// Set up routes
	app.SetupRoutes(r)

	// Start the Gin server
	fmt.Println("Starting server on port 8080")
	r.Run(":8080")
}
