package service_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"posts/internal/post"
	pb "posts/internal/proto"
	"posts/internal/service"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPosts(t *testing.T) {
	db := MockPostsDatabase{}
	s := service.NewPostsService(&db)

	_, err := s.GetPost(context.Background(), &pb.PostIdRequest{PostId: 13, AuthorId: 42})
	assert.Error(t, err)
	_, err = s.UpdatePost(context.Background(), &pb.UpdateRequest{PostId: 13, AuthorId: 42, Text: ""})
	assert.Error(t, err)

	req := &pb.CreateRequest{
		AuthorId: 42,
		Text:     "Sample post",
	}
	res, err := s.CreatePost(context.Background(), req)
	assert.NoError(t, err)
	postId := res.PostId

	getResp, err := s.GetPost(context.Background(), &pb.PostIdRequest{PostId: postId, AuthorId: 42})
	assert.NoError(t, err)
	assert.Equal(t, req.Text, getResp.Text)

	reqUpd := &pb.UpdateRequest{
		AuthorId: 42,
		PostId:   postId,
		Text:     "Updated text",
	}
	_, err = s.UpdatePost(context.Background(), reqUpd)
	assert.NoError(t, err)

	getResp, err = s.GetPost(context.Background(), &pb.PostIdRequest{PostId: postId, AuthorId: 42})
	assert.NoError(t, err)
	assert.Equal(t, reqUpd.Text, getResp.Text)

	_, err = s.DeletePost(context.Background(), &pb.PostIdRequest{PostId: postId, AuthorId: 42})
	assert.NoError(t, err)
	_, err = s.GetPost(context.Background(), &pb.PostIdRequest{PostId: postId, AuthorId: 42})
	assert.Error(t, err)
}

func TestGetPageOfPosts(t *testing.T) {
	db := MockPostsDatabase{}
	s := service.NewPostsService(&db)

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

	res, err := s.GetPageOfPosts(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, req.PageNum, res.PageNum)
	assert.Equal(t, int32(len(*posts)), res.PageSize)
	assert.Len(t, res.Posts, len(*posts))
	db.AssertExpectations(t)
}

type postExt struct {
	PostId      int32
	AuthorId    int32
	Text        string
	TimeUpdated string
}

type MockPostsDatabase struct {
	mock.Mock
	posts []postExt
}

func (m *MockPostsDatabase) CreatePost(authorId int32, text string) (int32, error) {
	postId := int32(len(m.posts) + 1)
	m.posts = append(m.posts,
		postExt{
			PostId:      postId,
			AuthorId:    authorId,
			Text:        text,
			TimeUpdated: time.Now().Format(time.RFC3339),
		})
	return postId, nil
}

func (m *MockPostsDatabase) UpdatePost(authorId, postId int32, text string) error {
	for i, p := range m.posts {
		if p.PostId == postId && p.AuthorId == authorId {
			m.posts[i] = postExt{
				PostId:      postId,
				AuthorId:    authorId,
				Text:        text,
				TimeUpdated: time.Now().Format(time.RFC3339),
			}
			return nil
		}
	}
	return fmt.Errorf("Not found")
}

func (m *MockPostsDatabase) DeletePost(authorId, postId int32) error {
	toDelete := -1
	for i, p := range m.posts {
		if p.PostId == postId && p.AuthorId == authorId {
			toDelete = i
			break
		}
	}
	if toDelete == -1 {
		return fmt.Errorf("Not found")
	}

	if len(m.posts) == 1 {
		m.posts = []postExt{}
		return nil
	}

	m.posts[toDelete] = m.posts[len(m.posts)-1]
	m.posts = m.posts[:1]
	return nil
}

func (m *MockPostsDatabase) GetPost(authorId, postId int32) (*post.Post, error) {
	for _, p := range m.posts {
		if p.PostId == postId && p.AuthorId == authorId {
			return &post.Post{Txt: p.Text, TimeUpdated: p.TimeUpdated}, nil
		}
	}

	return nil, fmt.Errorf("Not found")
}

func (m *MockPostsDatabase) GetPageOfPosts(Id, pageNum, pageSize int32) (*[]post.Post, error) {
	args := m.Called(Id, pageNum, pageSize)
	return args.Get(0).(*[]post.Post), args.Error(1)
}
