package model

import (
	"time"
)

// User represents a user in the system
type Admin struct {
	ID        int       `gorm:"primary_key;auto_increment"`
	OwnerName string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName sets the table name for the User model
func (Admin) TableName() string {
	return "users"
}
