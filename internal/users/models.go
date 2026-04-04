package users

import "time"

type User struct {
	ID           string    `db:"id" json:"id"`
	UUID         string    `db:"uuid" json:"uuid"`
	Username     string    `db:"username" json:"username"`
	Role         string    `db:"role" json:"role"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"password_hash"`
	Pfp          string    `db:"pfp" json:"pfp"`
	Description  string    `db:"description" json:"description"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
