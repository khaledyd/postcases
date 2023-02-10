package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// ConnectToPostgres creates a new connection to the PostgreSQL database.
func ConnectToPostgres() (*gorm.DB, error) {
	// Load database configuration
	// ...

	// Connect to the PostgreSQL database
	db, err := gorm.Open("postgres", "postgres://khaliddahir0200:ZhELiFRDuw62@ep-rough-mountain-919451.us-east-2.aws.neon.tech/neondb?options=project%3Dep-rough-mountain-919451")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	} else {
		fmt.Print("connected the dp")
	}
	// Configure the connection
	// ...

	return db, nil
}
