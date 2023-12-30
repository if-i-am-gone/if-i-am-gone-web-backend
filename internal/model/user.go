package model

import "time"

// User - holds all data required to represent a user
type User struct {
	ID         int       `json:"id,omitempty" db:"id"`
	FirstName  string    `json:"first_name,omitempty" db:"first_name"`
	LastName   string    `json:"last_name,omitempty" db:"last_name"`
	MiddleName string    `json:"middle_name,omitempty" db:"middle_name"`
	Email      string    `json:"email,omitempty" db:"email"`
	Hash       string    `json:"hash,omitempty" db:"hash"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
