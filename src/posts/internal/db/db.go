package db

import (
	"posts/internal/post"

	"github.com/jmoiron/sqlx"
)

type PostsDatabase interface {
	CreatePost(userId int64, text string) (int64, error)
	UpdatePost(userId, postId int64, text string) error
	DeletePost(userId, postId int64) error
	GetPost(userId, postId int64) (*post.Post, error)
	GetPageOfPosts(userId int64, pageNum, pageSize int32) (*[]post.Post, error)
}

type Database struct {
	PostsDatabase
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{PostsDatabase: NewPostsPostgres(db)}
}
