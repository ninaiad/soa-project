package service

import (
	"gateway/internal/db"
	posts_pb "gateway/internal/service/posts"
	stat_pb "gateway/internal/service/statistics"
	"gateway/internal/user"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
)

type Authorization interface {
	CreateUser(user user.User) (int64, error)
	UpdateUser(userId int64, update user.UserPublic) (user.UserPublic, error)
	GetUsername(userId int64) (string, error)
	GenerateToken(username, password string) (string, int64, error)
	ParseToken(token string) (int64, error)
}

type Statistics interface {
	AddEvent(postId, authorId, actorId int64, eventType EventType) error
}

type Service struct {
	Authorization
	Statistics
	posts_pb.PostsServerClient
	stat_pb.StatisticsServiceClient
}

func NewService(
	db *db.Database,
	postsCConn, statCConn grpc.ClientConnInterface,
	p *kafka.Producer, cfg KafkaConfig, ch chan kafka.Event) *Service {
	return &Service{
		Authorization:           CreateAuthService(db.Authorization),
		PostsServerClient:       posts_pb.NewPostsServerClient(postsCConn),
		StatisticsServiceClient: stat_pb.NewStatisticsServiceClient(statCConn),
		Statistics:              CreateKafkaService(p, cfg, ch),
	}
}
