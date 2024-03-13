package database

import (
	"mainservice"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user mainservice.User) error
	GetUser(username, password string) (mainservice.User, error)
	GetUserData(userId int) (mainservice.UserPublic, error)
	UpdateUser(userId int, update mainservice.UserPublic, timeUpdated string) error
}

type Database struct {
	Authorization
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
	}
}
