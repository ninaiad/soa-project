package statistics_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"soa-statistics/internal/common"
	pb "soa-statistics/internal/proto"
	"soa-statistics/internal/statistics"
)

func TestGetPostStatistics(t *testing.T) {
	db := MockStatisticsDatabase{}
	svc := statistics.NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.PostId{PostId: 1}
	postStats := &common.PostStatistics{
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
	db := MockStatisticsDatabase{}
	svc := statistics.NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_LIKE}
	postsStats := []common.PostStatistics{
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
	db := MockStatisticsDatabase{}
	svc := statistics.NewStatisticsService(&db)

	ctx := context.Background()
	req := &pb.TopKRequest{K: 2, Event: pb.EventType_VIEW}
	usersStats := []common.UserStatistics{
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

type MockStatisticsDatabase struct {
	mock.Mock
}

func (m *MockStatisticsDatabase) GetPostStatistics(ctx context.Context, postId uint64) (*common.PostStatistics, error) {
	args := m.Called(ctx, postId)
	return args.Get(0).(*common.PostStatistics), args.Error(1)
}

func (m *MockStatisticsDatabase) GetTopKPosts(ctx context.Context, eventType string, k uint64) ([]common.PostStatistics, error) {
	args := m.Called(ctx, eventType, k)
	return args.Get(0).([]common.PostStatistics), args.Error(1)
}

func (m *MockStatisticsDatabase) GetTopKUsers(ctx context.Context, eventType string, k uint64) ([]common.UserStatistics, error) {
	args := m.Called(ctx, eventType, k)
	return args.Get(0).([]common.UserStatistics), args.Error(1)
}
