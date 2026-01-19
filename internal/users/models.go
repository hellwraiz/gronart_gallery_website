package users

import "time"

type User struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Never send to frontend
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Optional: Session table (if using stateful tokens)
type Session struct {
	Token     string
	UserID    string
	ExpiresAt time.Time
	CreatedAt time.Time
}
