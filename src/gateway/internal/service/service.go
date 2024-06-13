package service

import (
	"gateway/internal/db"
	posts_pb "gateway/internal/service/posts_proto"
	stat_pb "gateway/internal/service/statistics_proto"
	"gateway/internal/user"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
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
