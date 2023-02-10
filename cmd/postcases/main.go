package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"postcases/internal/database"
)

func main() {
	// Create a new Gin instance
	r := gin.Default()
	r.SetTrustedProxies([]string{"	127.0.0.1"})

	fmt.Print("hi")

	// Load configuration from config package

	// Connect to the PostgreSQL database
	db, err := database.ConnectToPostgres()
	if err != nil {
		fmt.Println("failed to connect to PostgreSQL:", err)
		return
	}
	defer db.Close()

	// Connect to Redis

	// Connect to MongoDB

	// Load middleware

	// Load routes

	// Start the Gin server

	r.Run(":8080")
}
