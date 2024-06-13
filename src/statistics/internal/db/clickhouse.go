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

const createTableQueue = `CREATE TABLE IF NOT EXISTS queue (
    post UInt64,
    author UInt64,
	event Enum('like', 'view'),
    timestamp UInt64,
) ENGINE = Kafka('%s', '%s', '%s', '%s');`

const createTableEvents = `CREATE TABLE IF NOT EXISTS events (
    day Date,
    post UInt64,
    author UInt64,
	event Enum('like', 'view'),
	total UInt64
) ENGINE = SummingMergeTree(total) ORDER BY (day, post, author, event);`

const createView = `CREATE MATERIALIZED VIEW IF NOT EXISTS consumer TO events
	AS SELECT toDate(toDateTime(timestamp)) AS day, post, author, event, count() as total
	FROM queue GROUP BY day, post, author, event;`

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
	ctx context.Context, postId uint64) (*statistics.Post, error) {
	var dbPost statistics.Post
	query := fmt.Sprintf(
		`SELECT
			post,
			author,
			sum(multiIf(event = 'like', total, 0)) AS total_likes,
			sum(multiIf(event = 'view', total, 0)) AS total_views
		FROM events
		WHERE post = %v
		GROUP BY post, author`,
		postId)

	if err := db.conn.QueryRow(ctx, query).ScanStruct(&dbPost); err != nil {
		return nil, err
	}

	return &dbPost, nil
}

func (db *StatisticsClickhouse) GetTopKPosts(
	ctx context.Context, eventType string, k uint64) ([]statistics.Post, error) {
	var posts []statistics.Post
	query := fmt.Sprintf(
		`SELECT
			post,
			author,
			sum(multiIf(event = 'like', total, 0)) AS total_likes,
			sum(multiIf(event = 'view', total, 0)) AS total_views
		FROM events
		GROUP BY post, author
		ORDER BY multiIf('%v' = 'like', total_likes, total_views) DESC
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
	query := fmt.Sprintf(
		`SELECT
			author,
			sum(multiIf(event = 'like', total, 0)) AS total_likes,
			sum(multiIf(event = 'view', total, 0)) AS total_views
		FROM events
		GROUP BY author
		ORDER BY multiIf('%v' = 'like', total_likes, total_views) DESC
		LIMIT %v`,
		eventType, k)

	if err := db.conn.Select(ctx, &users, query); err != nil {
		return nil, err
	}

	return users, nil
}
