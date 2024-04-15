package service

import (
	"google.golang.org/grpc"
	"soa/main_service"
	"soa/main_service/pkg/database"
	pb "soa/posts"
)

type Authorization interface {
	CreateUser(user main_service.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	UpdateUser(userId int, update main_service.UserPublic) (main_service.UserPublic, error)
}

type Service struct {
	Authorization
	pb.PostsServerClient
}

func NewService(db *database.Database, cc grpc.ClientConnInterface) *Service {
	return &Service{
		Authorization:     CreateAuthService(db.Authorization),
		PostsServerClient: pb.NewPostsServerClient(cc),
	}
}
