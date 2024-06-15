package service_test

import (
	"context"
	"testing"

	pb "statistics/internal/pb"
	. "statistics/internal/service"
	"statistics/internal/statistics"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetPostStatistics(t *testing.T) {
	db := MockStatisticsDB{}
	svc := NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.PostId{PostId: 1}
	postStats := &statistics.Post{
		PostId:   1,
		AuthorId: 1,
		NumLikes: 10,
		NumViews: 100,
	}
	db.On("GetPostStatistics", ctx, req.GetPostId()).Return(postStats, nil)

	res, err := svc.GetPostStatistics(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, postStats.PostId, res.PostId)
	assert.Equal(t, postStats.AuthorId, res.AuthorId)
	assert.Equal(t, postStats.NumLikes, res.NumLikes)
	assert.Equal(t, postStats.NumViews, res.NumViews)
	db.AssertExpectations(t)
}

func TestGetTopKPosts(t *testing.T) {
	db := MockStatisticsDB{}
	svc := NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_LIKE}
	postsStats := []statistics.Post{
		{PostId: 1, AuthorId: 1, NumLikes: 10, NumViews: 100},
		{PostId: 2, AuthorId: 2, NumLikes: 15, NumViews: 150},
	}
	db.On("GetTopKPosts", ctx, "like", req.GetK()).Return(postsStats, nil)

	res, err := svc.GetTopKPosts(ctx, req)
	assert.NoError(t, err)
	assert.Len(t, res.Posts, len(postsStats))
	for i, post := range postsStats {
		assert.Equal(t, post.PostId, res.Posts[i].PostId)
		assert.Equal(t, post.AuthorId, res.Posts[i].AuthorId)
		assert.Equal(t, post.NumLikes, res.Posts[i].NumLikes)
		assert.Equal(t, post.NumViews, res.Posts[i].NumViews)
	}
	db.AssertExpectations(t)
}

func TestGetTopKUsers(t *testing.T) {
	db := MockStatisticsDB{}
	svc := NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_VIEW}
	usersStats := []statistics.User{
		{Id: 1, NumLikes: 10, NumViews: 100},
		{Id: 2, NumLikes: 15, NumViews: 150},
	}
	db.On("GetTopKUsers", ctx, "view", req.GetK()).Return(usersStats, nil)

	res, err := svc.GetTopKUsers(ctx, req)
	assert.NoError(t, err)
	assert.Len(t, res.Users, len(usersStats))
	for i, user := range usersStats {
		assert.Equal(t, user.Id, res.Users[i].Id)
		assert.Equal(t, user.NumLikes, res.Users[i].NumLikes)
		assert.Equal(t, user.NumViews, res.Users[i].NumViews)
	}
	db.AssertExpectations(t)
}

type MockStatisticsDB struct {
	mock.Mock
}

func (m *MockStatisticsDB) GetPostStatistics(
	ctx context.Context, postId int64) (*statistics.Post, error) {
	args := m.Called(ctx, postId)
	return args.Get(0).(*statistics.Post), args.Error(1)
}

func (m *MockStatisticsDB) GetTopKPosts(
	ctx context.Context, eventType string, k uint64) ([]statistics.Post, error) {
	args := m.Called(ctx, eventType, k)
	return args.Get(0).([]statistics.Post), args.Error(1)
}

func (m *MockStatisticsDB) GetTopKUsers(
	ctx context.Context, eventType string, k uint64) ([]statistics.User, error) {
	args := m.Called(ctx, eventType, k)
	return args.Get(0).([]statistics.User), args.Error(1)
}
