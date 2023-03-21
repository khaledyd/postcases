package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	FullName  string    `gorm:"not null" json:"full_name"`
	Email     string    `gorm:"not null;unique_index" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
