package service

import (
	"google.golang.org/grpc"
	"soa-main/internal/database"
	pb "soa-main/internal/posts_proto"
	"soa-main/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) error
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	UpdateUser(userId int, update user.UserPublic) (user.UserPublic, error)
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
