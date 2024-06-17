package db

import (
	"fmt"
	"log"
	"time"

	"posts/internal/post"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	postsTable = "posts"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func runDBMigration(cfg Config) error {
	m, err := migrate.New("file://./internal/db/migration",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	if err = runDBMigration(cfg); err != nil {
		return nil, err
	}

	log.Println("database is up and the migrations complete.")
	return db, nil
}

type PostsPostgres struct {
	db *sqlx.DB
}

func NewPostsPostgres(db *sqlx.DB) *PostsPostgres {
	return &PostsPostgres{db: db}
}

func (p *PostsPostgres) CreatePost(userId int64, text string) (int64, error) {
	var postId int64
	q := fmt.Sprintf(`
		INSERT INTO %s (time_updated, user_id, txt) values ($1, $2, $3) 
		RETURNING id`,
		postsTable)
	row := p.db.QueryRow(q, time.Now().Format(time.RFC3339), userId, text)
	err := row.Scan(&postId)
	return postId, err
}

func (p *PostsPostgres) UpdatePost(userId, postId int64, text string) error {
	q := fmt.Sprintf(`
		UPDATE %s SET time_updated=$1, txt=$2
		WHERE user_id=$3 AND id=$4`,
		postsTable)
	_, err := p.db.Exec(q, time.Now().Format(time.RFC3339), text, userId, postId)
	return err
}

func (p *PostsPostgres) DeletePost(userId, postId int64) error {
	q := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1 AND id=$2`,
		postsTable)
	_, err := p.db.Exec(q, userId, postId)
	return err
}

func (p *PostsPostgres) GetPost(userId, postId int64) (*post.Post, error) {
	var post post.Post
	q := fmt.Sprintf(`
		SELECT id, txt as text, time_updated as timeUpdated
		FROM %s
		WHERE user_id=$1 AND id=$2`,
		postsTable)
	err := p.db.Get(&post, q, userId, postId)
	return &post, err
}

func (p *PostsPostgres) GetPageOfPosts(
	userId int64, pageNum, pageSize int32) (*[]post.Post, error) {
	posts := []post.Post{}
	q := fmt.Sprintf(`
		SELECT s.*
		FROM (
			SELECT id, txt as text, time_updated as timeUpdated
			FROM %s
			WHERE user_id=$1
		) s 
		ORDER BY s.timeUpdated DESC 
		OFFSET $2
		LIMIT $3`,
		postsTable)
	err := p.db.Select(&posts, q, userId, (pageNum-1)*pageSize, pageSize)
	return &posts, err
}

func (p *PostsPostgres) DeleteUser(userId int64) error {
	q := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id=$1`,
		postsTable)
	_, err := p.db.Exec(q, userId)
	return err
}
