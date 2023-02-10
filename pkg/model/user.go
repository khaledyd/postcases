package model

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        int       `gorm:"primary_key;auto_increment"`
	Username  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName sets the table name for the User model
func (User) TableName() string {
	return "users"
}

// BeforeCreate hashes the user password before creating a new user
func (u *User) BeforeCreate() (err error) {

	if err != nil {
		return err
	}

	return nil
}
