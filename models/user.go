package models

import (
	"errors"

	"example.com/events_api/db"
	"example.com/events_api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hasedPassword, err := utils.GenerateHashword(u.Password)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(&u.Email, hasedPassword)
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, &u.Email)
	var dbPass string
	err := row.Scan(&u.Id, &dbPass)
	if err != nil {
		return err
	}
	isPasswordValid := utils.ComparePasswords(u.Password, dbPass)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}
	return nil
}
