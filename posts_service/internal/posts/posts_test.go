package posts_test

import (
	"soa-posts/internal/post"
	"soa-posts/internal/posts"
	"testing"
	"time"

	"context"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"google.golang.org/protobuf/types/known/emptypb"
	pb "soa-posts/internal/proto"
)

func TestCreatePost(t *testing.T) {
	db := MockPostsDatabase{}
	svc := posts.NewPostsService(&db)

	req := &pb.CreateRequest{
		AuthorId: 1,
		Text:     "Sample post",
	}
	db.On("CreatePost", req.AuthorId, req.Text).Return(13, nil)

	res, err := svc.CreatePost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, int32(13), res.PostId)
	db.AssertExpectations(t)
}

func TestUpdatePost(t *testing.T) {
	db := MockPostsDatabase{}
	svc := posts.NewPostsService(&db)

	req := &pb.UpdateRequest{
		AuthorId: 1,
		PostId:   23,
		Text:     "Updated text",
	}
	db.On("UpdatePost", req.AuthorId, req.PostId, req.Text).Return(nil)

	res, err := svc.UpdatePost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, &emptypb.Empty{}, res)
	db.AssertExpectations(t)
}

func TestDeletePost(t *testing.T) {
	db := MockPostsDatabase{}
	svc := posts.NewPostsService(&db)

	req := &pb.PostIdRequest{
		AuthorId: 2,
		PostId:   42,
	}
	db.On("DeletePost", req.AuthorId, req.PostId).Return(nil)

	res, err := svc.DeletePost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, &emptypb.Empty{}, res)
	db.AssertExpectations(t)
}

func TestGetPost(t *testing.T) {
	db := MockPostsDatabase{}
	svc := posts.NewPostsService(&db)

	req := &pb.PostIdRequest{
		AuthorId: 13,
		PostId:   23,
	}
	post := &post.Post{
		Txt:         "Sample post",
		TimeUpdated: time.Now().Format(time.RFC3339),
	}
	db.On("GetPost", req.AuthorId, req.PostId).Return(post, nil)

	res, err := svc.GetPost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, post.Txt, res.Text)
	assert.WithinDuration(t, time.Now(), res.TimeUpdated.AsTime(), time.Second)
	db.AssertExpectations(t)
}

func TestGetPageOfPosts(t *testing.T) {
	db := MockPostsDatabase{}
	svc := posts.NewPostsService(&db)

	req := &pb.GetPageOfPostsRequest{
		AuthorId: 1,
		PageNum:  1,
		PageSize: 2,
	}
	posts := &[]post.Post{
		{Txt: "Post 1", TimeUpdated: time.Now().Format(time.RFC3339)},
		{Txt: "Post 2", TimeUpdated: time.Now().Format(time.RFC3339)},
	}
	db.On("GetPageOfPosts", req.AuthorId, req.PageNum, req.PageSize).Return(posts, nil)

	res, err := svc.GetPageOfPosts(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.PageNum, res.PageNum)
	assert.Equal(t, int32(len(*posts)), res.PageSize)
	assert.Len(t, res.Posts, len(*posts))
	db.AssertExpectations(t)
}

type MockPostsDatabase struct {
	mock.Mock
}

func (m *MockPostsDatabase) CreatePost(authorId int32, text string) (int32, error) {
	args := m.Called(authorId, text)
	return int32(args.Int(0)), args.Error(1)
}

func (m *MockPostsDatabase) UpdatePost(authorId, postId int32, text string) error {
	args := m.Called(authorId, postId, text)
	return args.Error(0)
}

func (m *MockPostsDatabase) DeletePost(authorId, postId int32) error {
	args := m.Called(authorId, postId)
	return args.Error(0)
}

func (m *MockPostsDatabase) GetPost(authorId, postId int32) (*post.Post, error) {
	args := m.Called(authorId, postId)
	return args.Get(0).(*post.Post), args.Error(1)
}

func (m *MockPostsDatabase) GetPageOfPosts(authorId, pageNum, pageSize int32) (*[]post.Post, error) {
	args := m.Called(authorId, pageNum, pageSize)
	return args.Get(0).(*[]post.Post), args.Error(1)
}
