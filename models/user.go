package models

import "event-booking-api/db"

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	_, err := db.DB.Exec(query, u.Email, u.Password)

	return err
}
