package db

import (
	"posts/internal/post"

	"github.com/jmoiron/sqlx"
)

type PostsDatabase interface {
	CreatePost(userId int32, text string) (int32, error)
	UpdatePost(userId, postId int32, text string) error
	DeletePost(userId, postId int32) error
	GetPost(userId, postId int32) (*post.Post, error)
	GetPageOfPosts(userId, pageNum, pageSize int32) (*[]post.Post, error)
}

type Database struct {
	PostsDatabase
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{PostsDatabase: NewPostsPostgres(db)}
}
