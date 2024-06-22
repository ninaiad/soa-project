package service

import (
	"context"
	"time"

	"posts/internal/db"
	pb "posts/internal/pb"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostsService struct {
	pb.UnimplementedPostsServerServer
	db db.PostsDatabase
}

func NewPostsService(db db.PostsDatabase) *PostsService {
	return &PostsService{db: db}
}

func (p *PostsService) CreatePost(
	_ context.Context, r *pb.CreateRequest) (*pb.PostId, error) {
	postId, err := p.db.CreatePost(r.AuthorId, r.Text)
	if err != nil {
		return nil, err
	}

	return &pb.PostId{Id: postId}, nil
}

func (p *PostsService) UpdatePost(_ context.Context, r *pb.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.UpdatePost(r.AuthorId, r.PostId, r.Text)
}

func (p *PostsService) DeletePost(_ context.Context, r *pb.AuthoredPostId) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.DeletePost(r.AuthorId, r.PostId)
}

func (p *PostsService) GetPost(_ context.Context, r *pb.AuthoredPostId) (*pb.Post, error) {
	post, err := p.db.GetPost(r.AuthorId, r.PostId)
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, post.TimeUpdated)
	return &pb.Post{Id: post.Id, Text: post.Text, TimeUpdated: timestamppb.New(t)}, err
}

func (p *PostsService) GetPageOfPosts(
	_ context.Context, r *pb.PageOfPostsRequest) (*pb.PageOfPosts, error) {
	posts, err := p.db.GetPageOfPosts(r.AuthorId, r.PageNum, r.PageSize)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.Post, len(*posts))
	for i := range *posts {
		t, err := time.Parse(time.RFC3339, (*posts)[i].TimeUpdated)
		if err != nil {
			return nil, err
		}

		pbPosts[i] = &pb.Post{
			Id:          (*posts)[i].Id,
			Text:        (*posts)[i].Text,
			TimeUpdated: timestamppb.New(t),
		}
	}

	return &pb.PageOfPosts{
		PageNum:  r.PageNum,
		PageSize: int32(len(*posts)),
		AuthorId: r.AuthorId,
		Posts:    pbPosts,
	}, nil
}

func (p *PostsService) DeleteUser(ctx context.Context, in *pb.UserId) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.DeleteUser(in.GetId())
}
