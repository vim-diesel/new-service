package user

import (
	"net/mail"
	"time"
)

// User represents information about an individual user.
type User struct {
	ID        string
	FullName  string
	FirstName string
	LastName  string
	Email     mail.Address
	Enabled   bool
	CreatedAt time.Time
}

// NewUser contains information needed to create a new user.
type NewUser struct {
	FullName  string
	FirstName string
	LastName  string
	Email     mail.Address
	SubjectID string
}

// UpdateUser contains information needed to update a user.
type UpdateUser struct {
	Email     *mail.Address
	Enabled   *bool
	FullName  *string
	FirstName *string
	LastName  *string
}
