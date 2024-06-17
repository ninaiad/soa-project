package db

import (
	"context"
	"fmt"

	"statistics/internal/kafka"
	"statistics/internal/statistics"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Config struct {
	Host        string
	Port        int
	Username    string
	Password    string
	DBName      string
	TableFormat string
	KafkaCfg    kafka.Config
}

const createTableQueue = `
	CREATE TABLE IF NOT EXISTS queue (
		post_id Int64,
    author_id Int64,
		actor_id Int64,
	  event Enum('like', 'view'),
    timestamp UInt64,
	) ENGINE = Kafka('%s', '%s', '%s', '%s');`

const createTableEvents = `
	CREATE TABLE IF NOT EXISTS events (
		post_id Int64,
		author_id Int64,
		actor_id Int64,
		event Enum('like', 'view'),
		timestamp UInt64,
	) ENGINE = ReplacingMergeTree
	ORDER BY (post_id, author_id, actor_id, event);`

const createView = `
	CREATE MATERIALIZED VIEW IF NOT EXISTS consumer TO events AS 
		SELECT post_id, author_id, actor_id, event, timestamp
		FROM queue;`

func NewClickhouseDB(cfg Config) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)},
		Auth: clickhouse.Auth{
			Database: cfg.DBName,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
	})
	if err != nil {
		return nil, err
	}

	if err = conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	query := fmt.Sprintf(createTableQueue,
		cfg.KafkaCfg.KafkaAddr,
		cfg.KafkaCfg.KafkaTopic,
		cfg.KafkaCfg.KafkaGroupName,
		cfg.TableFormat)

	if err = conn.Exec(context.Background(), query); err != nil {
		return nil, err
	}
	if err = conn.Exec(context.Background(), createTableEvents); err != nil {
		return nil, err
	}
	return conn, conn.Exec(context.Background(), createView)
}

type StatisticsClickhouse struct {
	conn driver.Conn
}

func (db *StatisticsClickhouse) GetPostStatistics(
	ctx context.Context, postId int64) (*statistics.Post, error) {
	var dbPost statistics.Post
	query := fmt.Sprintf(`
		SELECT
			post_id,
			author_id,
			uniqExact(if(event = 'like', actor_id, NULL)) AS num_likes,
			uniqExact(if(event = 'view', actor_id, NULL)) AS num_views
		FROM events
		WHERE post_id = %v
		GROUP BY post_id, author_id`,
		postId)

	if err := db.conn.QueryRow(ctx, query).ScanStruct(&dbPost); err != nil {
		return nil, err
	}

	return &dbPost, nil
}

func (db *StatisticsClickhouse) GetTopKPosts(
	ctx context.Context, eventType string, k uint64) ([]statistics.Post, error) {
	var posts []statistics.Post
	query := fmt.Sprintf(`
		SELECT
			post_id,
			author_id,
			uniqExact(if(event = 'like', actor_id, NULL)) AS num_likes,
			uniqExact(if(event = 'view', actor_id, NULL)) AS num_views 
		FROM events
		GROUP BY post_id, author_id
		ORDER BY if('%v' = 'like', num_likes, num_views) DESC
		LIMIT %v`,
		eventType, k)

	if err := db.conn.Select(ctx, &posts, query); err != nil {
		return nil, err
	}

	return posts, nil
}

func (db *StatisticsClickhouse) GetTopKUsers(
	ctx context.Context, eventType string, k uint64) ([]statistics.User, error) {
	var users []statistics.User
	query := fmt.Sprintf(`
		SELECT
			author_id,
			sum(num_likes) AS num_likes,
			sum(num_views) AS num_views
		FROM (
			SELECT
				author_id,
				if(event = 'like', uniqExact(post_id, actor_id), 0) AS num_likes,
				if(event = 'view', uniqExact(post_id, actor_id), 0) AS num_views
			FROM (
				SELECT author_id, post_id, actor_id, event
				FROM events
				GROUP BY author_id, actor_id, post_id, event
			)
			GROUP BY author_id, event
		)
		GROUP BY author_id
		ORDER BY if('%v' = 'like', num_likes, num_views) DESC
		LIMIT %v`,
		eventType, k)

	if err := db.conn.Select(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}

func (db *StatisticsClickhouse) DeleteUser(ctx context.Context, userId int64) error {
	query := fmt.Sprintf(`DELETE from events WHERE author_id = %v`, userId)
	return db.conn.Exec(ctx, query)
}

func (db *StatisticsClickhouse) DeletePost(ctx context.Context, postId int64) error {
	query := fmt.Sprintf(`DELETE from events WHERE post_id = %v`, postId)
	return db.conn.Exec(ctx, query)
}
