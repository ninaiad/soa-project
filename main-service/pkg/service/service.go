package service

import (
	"mainservice"
	"mainservice/pkg/database"
)

type Authorization interface {
	CreateUser(user mainservice.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	UpdateUser(userId int, update mainservice.UserPublic) (mainservice.UserPublic, error)
}

type Service struct {
	Authorization
}

func NewService(db *database.Database) *Service {
	return &Service{
		Authorization: CreateAuthService(db.Authorization),
	}
}
