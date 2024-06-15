package db

import (
	"gateway/internal/user"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user user.User) (int64, error)
	GetUserId(username, password string) (int64, error)
	GetUserData(userId int64) (user.UserPublic, error)
	UpdateUser(userId int64, update user.UserPublic, timeUpdated string) error
}

type Database struct {
	Authorization
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{
		Authorization: NewAuthPostgres(db),
	}
}
