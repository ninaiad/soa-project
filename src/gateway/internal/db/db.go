package db

import (
	"gateway/internal/user"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	CreateUser(user user.User) (int64, error)
	DeleteUser(userId int64) error
	GetUserId(username, password string) (int64, error)
	GetUserData(userId int64) (user.UserPublic, error)
	UpdateUser(userId int64, update user.UserPublic, timeUpdated string) error
}

func NewDatabase(db *sqlx.DB) Database {
	return NewAuthPostgres(db)
}
