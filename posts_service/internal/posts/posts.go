package posts

import (
	"context"
	"time"

	"soa-posts/internal/database"
	pb "soa-posts/internal/proto"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PostsService struct {
	pb.UnimplementedPostsServerServer
	db database.PostsDatabase
}

func NewPostsService(db database.PostsDatabase) *PostsService {
	return &PostsService{db: db}
}

func (p *PostsService) CreatePost(_ context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	postId, err := p.db.CreatePost(req.AuthorId, req.Text)
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{PostId: postId}, nil
}

func (p *PostsService) UpdatePost(_ context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.UpdatePost(req.AuthorId, req.PostId, req.Text)
}

func (p *PostsService) DeletePost(_ context.Context, req *pb.PostIdRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, p.db.DeletePost(req.AuthorId, req.PostId)
}

func (p *PostsService) GetPost(_ context.Context, req *pb.PostIdRequest) (*pb.Post, error) {
	post, err := p.db.GetPost(req.AuthorId, req.PostId)
	if err != nil {
		return nil, err
	}

	t, err := time.Parse(time.RFC3339, post.TimeUpdated)
	return &pb.Post{Text: post.Txt, TimeUpdated: timestamppb.New(t)}, err
}

func (p *PostsService) GetPageOfPosts(_ context.Context, req *pb.GetPageOfPostsRequest) (*pb.GetPageOfPostsResponse, error) {
	posts, err := p.db.GetPageOfPosts(req.AuthorId, req.PageNum, req.PageSize)
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

	return &pb.GetPageOfPostsResponse{PageNum: req.PageNum, PageSize: int32(len(*posts)), Posts: pbPosts}, nil
}
