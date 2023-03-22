package main

import (
	"context"
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

	// Connect to the MongoDB Atlas database

	client, err := database.ConnectMongoAtlas()
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return
	}
	defer client.Disconnect(context.Background())
	database.MongoClient = client

	// Create a new Gin instance
	r := gin.Default()

	// Set up routes
	app.SetupRoutes(r)

	// Start the Gin server
	fmt.Println("Starting server on port 8080")
	r.Run(":8080")
}
