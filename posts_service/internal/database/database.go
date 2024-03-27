package database

import (
	"soa/posts_service"

	"github.com/jmoiron/sqlx"
)

type PostsDatabase interface {
	CreatePost(userId int32, text string) (int32, error)
	UpdatePost(userId, postId int32, text string) error
	DeletePost(userId, postId int32) error
	GetPost(userId, postId int32) (*posts_service.Post, error)
	GetPageOfPosts(userId, pageNum, pageSize int32) (*[]posts_service.Post, error)
}

type Database struct {
	PostsDatabase
}

func NewDatabase(db *sqlx.DB) *Database {
	return &Database{PostsDatabase: NewPostsPostgres(db)}
}
