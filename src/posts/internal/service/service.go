package service

import (
	"context"
	"time"

	"posts/internal/db"
	pb "posts/internal/proto"

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
	_ context.Context, r *pb.CreateRequest) (*pb.CreateResponse, error) {
	postId, err := p.db.CreatePost(r.AuthorId, r.Text)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{PostId: postId}, nil
}

func (p *PostsService) UpdatePost(_ context.Context, r *pb.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.UpdatePost(r.AuthorId, r.PostId, r.Text)
}

func (p *PostsService) DeletePost(_ context.Context, r *pb.PostIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.DeletePost(r.AuthorId, r.PostId)
}

func (p *PostsService) GetPost(_ context.Context, r *pb.PostIdRequest) (*pb.Post, error) {
	post, err := p.db.GetPost(r.AuthorId, r.PostId)
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, post.TimeUpdated)
	return &pb.Post{Text: post.Txt, TimeUpdated: timestamppb.New(t)}, err
}

func (p *PostsService) GetPageOfPosts(
	_ context.Context, r *pb.GetPageOfPostsRequest) (*pb.GetPageOfPostsResponse, error) {
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

		pbPosts[i] = &pb.Post{Text: (*posts)[i].Txt, TimeUpdated: timestamppb.New(t)}
	}

	return &pb.GetPageOfPostsResponse{
		PageNum:  r.PageNum,
		PageSize: int32(len(*posts)),
		Posts:    pbPosts,
	}, nil
}