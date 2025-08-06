package service

import (
	"codebase-golang/pkg/database"
)

type User struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func FetchUsers() ([]User, error) {
	var users []User
	query := "SELECT id, name FROM users"

	err := database.DB.Select(&users, query)
	return users, err
}
