package database

import (
	"soa/main_service"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user main_service.User) error
	GetUser(username, password string) (main_service.User, error)
	GetUserData(userId int) (main_service.UserPublic, error)
	UpdateUser(userId int, update main_service.UserPublic, timeUpdated string) error
}

type Database struct {
	Authorization
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
	}
}
