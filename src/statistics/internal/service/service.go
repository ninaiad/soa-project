package service

import (
	"context"

	"statistics/internal/db"
	pb "statistics/internal/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type StatisticsService struct {
	pb.UnimplementedStatisticsServiceServer
	db db.StatisticsDatabase
}

func NewStatisticsService(db db.StatisticsDatabase) *StatisticsService {
	return &StatisticsService{db: db}
}

func (s *StatisticsService) GetPostStatistics(
	ctx context.Context, in *pb.PostId) (*pb.PostStatistics, error) {
	postDb, err := s.db.GetPostStatistics(ctx, in.GetPostId())
	if err != nil {
		return nil, err
	}

	return &pb.PostStatistics{
		PostId:   postDb.PostId,
		AuthorId: postDb.AuthorId,
		NumLikes: postDb.NumLikes,
		NumViews: postDb.NumViews,
	}, err
}

func (s *StatisticsService) GetTopKPosts(
	ctx context.Context, in *pb.TopKRequest) (*pb.TopPosts, error) {
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
				PostId:   p.PostId,
				AuthorId: p.AuthorId,
				NumLikes: p.NumLikes,
				NumViews: p.NumViews,
			})
	}

	return &pb.TopPosts{Posts: posts, TimeCollected: timestamppb.Now()}, nil
}

func (s *StatisticsService) GetTopKUsers(
	ctx context.Context, in *pb.TopKRequest) (*pb.TopUsers, error) {
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

	users := []*pb.UserStatistics{}
	for _, u := range usersDb {
		users = append(users,
			&pb.UserStatistics{
				Id:       u.Id,
				NumLikes: u.NumLikes,
				NumViews: u.NumViews,
			})
	}

	return &pb.TopUsers{Users: users, TimeCollected: timestamppb.Now()}, nil
}
