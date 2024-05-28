package statistics

import (
	"context"
	"soa-statistics/internal/database"
	pb "soa-statistics/internal/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatisticsService struct {
	pb.UnimplementedStatisticsServiceServer
	db database.StatisticsDatabase
}

func NewStatisticsService(db database.StatisticsDatabase) *StatisticsService {
	return &StatisticsService{db: db}
}

func (s *StatisticsService) GetPostStatistics(ctx context.Context, in *pb.PostId) (*pb.PostStatistics, error) {
	postDb, err := s.db.GetPostStatistics(ctx, uint64(in.GetPostId()))
	if err != nil {
		return nil, err
	}

	return &pb.PostStatistics{
		PostId:   int32(postDb.PostId),
		AuthorId: int32(postDb.AuthorId),
		NumLikes: postDb.TotalLikes,
		NumViews: postDb.TotalViews,
	}, err
}

func (s *StatisticsService) GetTopKPosts(ctx context.Context, in *pb.TopKRequest) (*pb.TopPosts, error) {
	var eventType string
	if in.GetEvent() == pb.EventType_LIKE {
		eventType = "like"
	} else {
		eventType = "view"
	}

	postsDb, err := s.db.GetTopKPosts(ctx, eventType, in.GetK())
	if err != nil {
		return nil, err
	}

	posts := []*pb.PostStatistics{}
	for _, p := range postsDb {
		posts = append(posts,
			&pb.PostStatistics{
				PostId:   int32(p.PostId),
				AuthorId: int32(p.AuthorId),
				NumLikes: p.TotalLikes,
				NumViews: p.TotalViews,
			})
	}

	return &pb.TopPosts{Posts: posts, TimeCollected: timestamppb.Now()}, nil
}

func (s *StatisticsService) GetTopKUsers(ctx context.Context, in *pb.TopKRequest) (*pb.TopUsers, error) {
	var eventType string
	if in.GetEvent() == pb.EventType_LIKE {
		eventType = "like"
	} else {
		eventType = "view"
	}

	usersDb, err := s.db.GetTopKUsers(ctx, eventType, in.GetK())
	if err != nil {
		return nil, err
	}

	users := []*pb.UsersStatistics{}
	for _, u := range usersDb {
		users = append(users,
			&pb.UsersStatistics{
				AuthorId: int32(u.AuthorId),
				NumLikes: u.TotalLikes,
				NumViews: u.TotalViews,
			})
	}

	return &pb.TopUsers{Users: users, TimeCollected: timestamppb.Now()}, nil
}
