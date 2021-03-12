package photohistory

import (
	"errors"

	"github.com/google/uuid"
)

// User of the system
type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Firstname string    `db:"first_name" json:"first_name"`
	Lastname  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Phone     string    `db:"phone" json:"phone"`
}

// Common Errors
var (
	ErrUserNotFound = errors.New("user not found")
)
