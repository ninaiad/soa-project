package service_test

import (
	"context"
	"testing"

	pb "statistics/internal/proto"
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
		PostId:     1,
		AuthorId:   1,
		TotalLikes: 10,
		TotalViews: 100,
	}
	db.On("GetPostStatistics", ctx, uint64(req.GetPostId())).Return(postStats, nil)

	res, err := svc.GetPostStatistics(ctx, req)
	assert.NoError(t, err)
	assert.Equal(t, int32(postStats.PostId), res.PostId)
	assert.Equal(t, int32(postStats.AuthorId), res.AuthorId)
	assert.Equal(t, postStats.TotalLikes, res.NumLikes)
	assert.Equal(t, postStats.TotalViews, res.NumViews)
	db.AssertExpectations(t)
}

func TestGetTopKPosts(t *testing.T) {
	db := MockStatisticsDB{}
	svc := NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_LIKE}
	postsStats := []statistics.Post{
		{PostId: 1, AuthorId: 1, TotalLikes: 10, TotalViews: 100},
		{PostId: 2, AuthorId: 2, TotalLikes: 15, TotalViews: 150},
	}
	db.On("GetTopKPosts", ctx, "like", req.GetK()).Return(postsStats, nil)

	res, err := svc.GetTopKPosts(ctx, req)
	assert.NoError(t, err)
	assert.Len(t, res.Posts, len(postsStats))
	for i, post := range postsStats {
		assert.Equal(t, int32(post.PostId), res.Posts[i].PostId)
		assert.Equal(t, int32(post.AuthorId), res.Posts[i].AuthorId)
		assert.Equal(t, post.TotalLikes, res.Posts[i].NumLikes)
		assert.Equal(t, post.TotalViews, res.Posts[i].NumViews)
	}
	db.AssertExpectations(t)
}

func TestGetTopKUsers(t *testing.T) {
	db := MockStatisticsDB{}
	svc := NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_VIEW}
	usersStats := []statistics.User{
		{AuthorId: 1, TotalLikes: 10, TotalViews: 100},
		{AuthorId: 2, TotalLikes: 15, TotalViews: 150},
	}
	db.On("GetTopKUsers", ctx, "view", req.GetK()).Return(usersStats, nil)

	res, err := svc.GetTopKUsers(ctx, req)
	assert.NoError(t, err)
	assert.Len(t, res.Users, len(usersStats))
	for i, user := range usersStats {
		assert.Equal(t, int32(user.AuthorId), res.Users[i].AuthorId)
		assert.Equal(t, user.TotalLikes, res.Users[i].NumLikes)
		assert.Equal(t, user.TotalViews, res.Users[i].NumViews)
	}
	db.AssertExpectations(t)
}

type MockStatisticsDB struct {
	mock.Mock
}

func (m *MockStatisticsDB) GetPostStatistics(
	ctx context.Context, postId uint64) (*statistics.Post, error) {
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
