package models

import "time"

// User - holds all data required to represent a user
type User struct {
	ID         int       `json:"id,omitempty"`
	FirstName  string    `json:"first_name,omitempty"`
	LastName   string    `json:"last_name,omitempty"`
	MiddleName string    `json:"middle_name,omitempty"`
	Email      string    `json:"email,omitempty"`
	Hash       string    `json:"hash,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}
