package models

import (
	"errors"

	"example.com/go-api-practice/db"
	"example.com/go-api-practice/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES(?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	hashedPassword, err := utils.HashData(u.Password)
	if err != nil {
		return err
	}
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	u.Password = hashedPassword

	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")

	}
	isValidPassword := utils.CheckDataHash(u.Password, retrievedPassword)
	if !isValidPassword {
		return errors.New("credentials invalid")
	}
	return nil
}
