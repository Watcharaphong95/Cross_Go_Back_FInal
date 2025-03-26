package dto

import "time"

type User struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Address     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
