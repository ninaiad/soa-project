package db

import (
	"context"

	"statistics/internal/statistics"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StatisticsDatabase interface {
	GetPostStatistics(ctx context.Context, postId int64) (*statistics.Post, error)
	GetTopKPosts(ctx context.Context, eventType string, k uint64) ([]statistics.Post, error)
	GetTopKUsers(ctx context.Context, eventType string, k uint64) ([]statistics.User, error)

	DeleteUser(ctx context.Context, userId int64) error
	DeletePost(ctx context.Context, postId int64) error
}

func NewDatabase(conn driver.Conn) StatisticsDatabase {
	return &StatisticsClickhouse{conn: conn}
}
