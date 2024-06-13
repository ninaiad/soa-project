package db

import (
	"gateway/internal/user"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GetUser(username, password string) (user.User, error)
	GetUserLogin(userId int) (user.User, error)
	GetUserData(userId int) (user.UserPublic, error)
	UpdateUser(userId int, update user.UserPublic, timeUpdated string) error
}

type Database struct {
	Authorization
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
	}
}
