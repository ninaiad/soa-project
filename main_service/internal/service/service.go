package service

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
	"soa-main/internal/database"
	posts_pb "soa-main/internal/posts_proto"
	stat_pb "soa-main/internal/statistics_proto"
	"soa-main/internal/user"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	UpdateUser(userId int, update user.UserPublic) (user.UserPublic, error)
	GetUserLogin(userId int) (user.User, error)
	GenerateToken(username, password string) (string, int, error)
	ParseToken(token string) (int, error)
}

type Statistics interface {
	AddEvent(postId int64, authorId int64, eventType EventType) error
}

type Service struct {
	Authorization
	Statistics
	posts_pb.PostsServerClient
	stat_pb.StatisticsServiceClient
}

func NewService(db *database.Database, posts_cc, stat_cc grpc.ClientConnInterface, kafkaProducer *kafka.Producer, cfg KafkaConfig, eventCh chan kafka.Event) *Service {
	return &Service{
		Authorization:           CreateAuthService(db.Authorization),
		PostsServerClient:       posts_pb.NewPostsServerClient(posts_cc),
		StatisticsServiceClient: stat_pb.NewStatisticsServiceClient(stat_cc),
		Statistics:              CreateKafkaService(kafkaProducer, cfg, eventCh),
	}
}
