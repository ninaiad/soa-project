package database

import (
	"fmt"
	"soa/posts_service"
	"time"

	"github.com/jmoiron/sqlx"
)

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (p *PostsPostgres) CreatePost(userId int32, text string) (int32, error) {
	var postId int32
	q := fmt.Sprintf("INSERT INTO %s (timeUpdated, user_id, txt) values ($1, $2, $3) RETURNING id", postsTable)
	row := p.db.QueryRow(q, time.Now().Format(time.RFC3339), userId, text)
	err := row.Scan(&postId)
	return postId, err
}

func (p *PostsPostgres) UpdatePost(userId, postId int32, text string) error {
	q := fmt.Sprintf("UPDATE %s SET timeUpdated=$1, txt=$2  WHERE user_id=$3 AND id=$4", postsTable)
	_, err := p.db.Exec(q, time.Now().Format(time.RFC3339), text, userId, postId)
	return err
}

func (p *PostsPostgres) DeletePost(userId, postId int32) error {
	q := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", postsTable)
	_, err := p.db.Exec(q, userId, postId)
	return err
}

func (p *PostsPostgres) GetPost(userId, postId int32) (*posts_service.Post, error) {
	var post posts_service.Post
	q := fmt.Sprintf("SELECT txt, timeUpdated FROM %s WHERE user_id=$1 AND id=$2", postsTable)
	err := p.db.Get(&post, q, userId, postId)
	return &post, err
}

func (p *PostsPostgres) GetPageOfPosts(userId, pageNum, pageSize int32) (*[]posts_service.Post, error) {
	posts := []posts_service.Post{}
	q := fmt.Sprintf("SELECT s.* FROM (SELECT txt, timeUpdated FROM %s WHERE user_id=$1) s ORDER BY s.timeUpdated DESC OFFSET $2 LIMIT $3", postsTable)
	err := p.db.Select(&posts, q, userId, (pageNum-1)*pageSize, pageSize)
	return &posts, err
}
