package service

import (
	"context"

	posts_pb "gateway/internal/service/posts_pb"
	stat_pb "gateway/internal/service/statistics_pb"
)

func (s *Service) CreatePost(authorId int64, text string) (int64, error) {
	r, err := s.pClient.CreatePost(context.Background(),
		&posts_pb.CreateRequest{AuthorId: authorId, Text: text})
	if err != nil {
		return 0, err
	}

	return r.GetId(), nil
}

func (s *Service) UpdatePost(authorId, postId int64, text string) error {
	_, err := s.pClient.UpdatePost(context.Background(),
		&posts_pb.UpdateRequest{AuthorId: authorId, PostId: postId, Text: text})
	return err
}

func (s *Service) DeletePost(authorId, postId int64) error {
	_, err := s.pClient.DeletePost(context.Background(),
		&posts_pb.AuthoredPostId{AuthorId: authorId, PostId: postId})
	if err != nil {
		return err
	}

	_, err = s.sClient.DeletePost(context.Background(), &stat_pb.PostId{Id: postId})
	return err
}

func (s *Service) GetPost(authorId, postId int64) (*posts_pb.Post, error) {
	return s.pClient.GetPost(context.Background(),
		&posts_pb.AuthoredPostId{AuthorId: authorId, PostId: postId})
}

func (s *Service) GetPageOfPosts(
	authorId int64, pageNum, pageSize int32) (*posts_pb.PageOfPosts, error) {
	return s.pClient.GetPageOfPosts(context.Background(),
		&posts_pb.PageOfPostsRequest{AuthorId: authorId, PageNum: pageNum, PageSize: pageSize})
}
