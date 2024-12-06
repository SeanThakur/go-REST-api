package models

import (
	"errors"
	"seanThakur/go-restapi/db"
	"seanThakur/go-restapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users(email, password) VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCreds() error {
	query := `SELECT id, password from users where email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var reterivedPassword string
	err := row.Scan(&u.ID, &reterivedPassword)

	if err != nil {
		return errors.New("credential invalid")
	}

	isValid := utils.CheckPasswordFromHash(u.Password, reterivedPassword)
	if !isValid {
		return errors.New("invalid password")
	}

	return nil
}
