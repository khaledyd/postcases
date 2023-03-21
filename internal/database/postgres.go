package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/khaledyd/postcases/pkg/model"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

// ConnectToPostgres creates a new connection to the PostgreSQL database.
func ConnectToPostgres() (*gorm.DB, error) {
	DB, err := gorm.Open("postgres", "postgres://root:cusbo@localhost:5432/postcases?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %v", err)
	}

	if err := DB.DB().Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	// Configure the connection
	DB.AutoMigrate(model.User{})

	fmt.Println("Connected to database")
	return DB, nil
}
