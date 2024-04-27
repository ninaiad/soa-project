package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
	"soa-main/internal/database"
	pb "soa-main/internal/posts_proto"
	"soa-main/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, int, error)
	ParseToken(token string) (int, error)
	UpdateUser(userId int, update user.UserPublic) (user.UserPublic, error)
}

type Statistics interface {
	AddEvent(postId int64, authorId int64, eventType EventType) error
}

type Service struct {
	Authorization
	Statistics
	pb.PostsServerClient
}

func NewService(db *database.Database, cc grpc.ClientConnInterface, kafkaProducer *kafka.Producer, cfg KafkaConfig, eventCh chan kafka.Event) *Service {
	return &Service{
		Authorization:     CreateAuthService(db.Authorization),
		PostsServerClient: pb.NewPostsServerClient(cc),
		Statistics:        CreateKafkaService(kafkaProducer, cfg, eventCh),
	}
}
