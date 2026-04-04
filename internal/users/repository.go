package users

import (
	"fmt"
	"gronart_gallery_website/internal/auth"
	"log"

	"github.com/jmoiron/sqlx"
)

func CreateUser(db *sqlx.DB, u *User) error {

	query := `INSERT INTO users (uuid, email, password_hash, pfp, username, description) VALUES (:uuid, :email, :password_hash, :pfp, :username, :description)`

	u.UUID = generateUUID()

	if result, err := db.NamedExec(query, u); err != nil {
		return fmt.Errorf("Failed to create user: %s", err)
	} else if numAffected, errResult := result.RowsAffected(); numAffected == 0 && errResult == nil {
		return fmt.Errorf("Failed to create user: Database unaffected")
	} else if errResult != nil {
		return fmt.Errorf("Couldn't find out if user was created: %s", errResult)
	}
	return nil
}

func DeleteUser(db *sqlx.DB, uuid string) error {

	query := `DELETE FROM users WHERE uuid = ?`

	if result, err := db.NamedExec(query, uuid); err != nil {
		return fmt.Errorf("Failed to delete user: %s", err)
	} else if numAffected, errResult := result.RowsAffected(); numAffected == 0 && errResult == nil {
		return fmt.Errorf("Failed to delete user: Database unaffected")
	} else if errResult != nil {
		return fmt.Errorf("Couldn't find out if user was deleted: %s", errResult)
	}
	return nil
}

func GetUserByUUID(db *sqlx.DB, uuid string) (*User, error) {
	var user User

	query := `
	SELECT * FROM users
	WHERE uuid = ?
	`

	if err := db.Get(&user, query, uuid); err != nil {
		return nil, fmt.Errorf("Failed to get user with token %s: %s", uuid, err)
	}
	return &user, nil
}

func Login(db *sqlx.DB, login *auth.Login) (string, error) {
	var user User

	query := `
	SELECT password_hash FROM users
	WHERE email = ?
	`

	if err := db.Get(&user, query, login.Email); err != nil {
		log.Print("email wrong")
		return "", fmt.Errorf("User couldn't be authenticated")
	}

	if !CheckPassword(login.Password, user.PasswordHash) {
		log.Print("password wrong")
		return "", fmt.Errorf("User couldn't be authenticated")
	}
	token, err := generateToken(user.UUID, user.Role)
	if err != nil {
		return "", fmt.Errorf("Couldn't generate token: %s", err)
	}
	return token, nil

}
