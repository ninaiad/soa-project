package service

import (
	"gateway/internal/db"
	posts_pb "gateway/internal/service/posts_pb"
	stat_pb "gateway/internal/service/statistics_pb"
	"gateway/internal/user"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"google.golang.org/grpc"
)

type IService interface {
	// Auth
	CreateUser(user user.User) (int64, error)
	UpdateUser(userId int64, update user.UserPublic) (user.UserPublic, error)
	DeleteUser(userId int64) error
	GetUsername(userId int64) (string, error)
	GenerateToken(username, password string) (string, int64, error)
	ParseToken(token string) (int64, error)

	// Posts
	CreatePost(authorId int64, text string) (int64, error)
	UpdatePost(authorId, postId int64, text string) error
	DeletePost(authorId, postId int64) error
	GetPost(authorId, postId int64) (*posts_pb.Post, error)
	GetPageOfPosts(authorId int64, pageNum, pageSize int32) (*posts_pb.PageOfPosts, error)

	// Statistics
	AddEvent(postId, authorId, actorId int64, eventType EventType) error
	GetPostStatistics(postId int64) (*stat_pb.PostStatistics, error)
	GetTopKPosts(event EventType, k uint64) (*stat_pb.TopPosts, error)
	GetTopKUsers(event EventType, k uint64) (*stat_pb.TopUsers, error)
}

type Service struct {
	db db.Database

	pClient posts_pb.PostsServerClient
	sClient stat_pb.StatisticsServiceClient

	kafkaProducer *kafka.Producer
	kafkaCfg      KafkaConfig
	kafkaEventCh  chan kafka.Event
}

func NewService(
	db db.Database,
	postsCConn, statCConn grpc.ClientConnInterface,
	p *kafka.Producer, cfg KafkaConfig, ch chan kafka.Event) IService {
	return &Service{
		db:            db,
		pClient:       posts_pb.NewPostsServerClient(postsCConn),
		sClient:       stat_pb.NewStatisticsServiceClient(statCConn),
		kafkaProducer: p,
		kafkaCfg:      cfg,
		kafkaEventCh:  ch,
	}
}
