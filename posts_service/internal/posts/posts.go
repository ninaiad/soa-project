package posts

import (
	"context"

	pb "soa/posts"
	"soa/posts_service/internal/database"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PostsService struct {
	pb.UnimplementedPostsServerServer
	db database.PostsDatabase
}

func NewPostsService() *PostsService {
	return &PostsService{}
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

	return &pb.Post{Text: post.Txt, TimeUpdated: &post.TimeUpdated}, nil
}

func (p *PostsService) GetPageOfPosts(_ context.Context, req *pb.GetPageOfPostsRequest) (*pb.GetPageOfPostsResponse, error) {
	posts, err := p.db.GetPageOfPosts(req.AuthorId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	pbPosts := make([]*pb.Post, len(*posts))
	for i := range *posts {
		pbPosts = append(pbPosts, &pb.Post{Text: (*posts)[i].Txt, TimeUpdated: &(*posts)[i].TimeUpdated})
	}

	return &pb.GetPageOfPostsResponse{PageNum: req.PageNum, PageSize: int32(len(*posts)), Posts: pbPosts}, nil
}	
