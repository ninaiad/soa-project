package db

import (
	"context"

	"statistics/internal/statistics"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StatisticsDatabase interface {
	GetPostStatistics(ctx context.Context, postId uint64) (*statistics.Post, error)
	GetTopKPosts(ctx context.Context, eventType string, k uint64) ([]statistics.Post, error)
	GetTopKUsers(ctx context.Context, eventType string, k uint64) ([]statistics.User, error)
}

func NewDatabase(conn driver.Conn) StatisticsDatabase {
	return &StatisticsClickhouse{conn: conn}
}
