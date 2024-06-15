package main

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "posts_tests/pb"
)

func TestPost(t *testing.T) {
	conn, err := grpc.NewClient(
		os.Getenv("POSTS_SERVER_ADDR"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.NoError(t, err)
	defer conn.Close()

	client := pb.NewPostsServerClient(conn)

	txt := "Test post"
	res, err := client.CreatePost(context.Background(), &pb.CreateRequest{AuthorId: 42, Text: txt})
	assert.NoError(t, err)
	postId := res.PostId

	req := &pb.PostIdRequest{AuthorId: 42, PostId: postId}

	resp, err := client.GetPost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Text, txt)

	_, err = client.UpdatePost(context.Background(),
		&pb.UpdateRequest{
			AuthorId: 42,
			PostId:   postId,
			Text:     "Updated"})
	assert.NoError(t, err)

	resp, err = client.GetPost(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, resp.Text, "Updated")

	_, err = client.DeletePost(context.Background(), req)
	assert.NoError(t, err)
	_, err = client.GetPost(context.Background(), req)
	assert.Error(t, err)

	ids := []int64{}
	for i := 0; i < 5; i++ {
		res, err = client.CreatePost(context.Background(),
			&pb.CreateRequest{AuthorId: 42, Text: fmt.Sprint(i)})
		assert.NoError(t, err)
		ids = append(ids, res.PostId)
	}

	page, err := client.GetPageOfPosts(context.Background(),
		&pb.GetPageOfPostsRequest{AuthorId: 42, PageNum: 1, PageSize: 6})
	assert.NoError(t, err)

	assert.Equal(t, page.PageNum, int32(1))
	assert.Equal(t, page.PageSize, int32(5))
	for _, p := range page.Posts {
		assert.Contains(t, []string{"0", "1", "2", "3", "4"}, p.Text)
	}

	for _, id := range ids {
		_, err = client.DeletePost(context.Background(), &pb.PostIdRequest{AuthorId: 42, PostId: id})
		assert.NoError(t, err)
	}
}
