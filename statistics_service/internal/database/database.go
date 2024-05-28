package database

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"soa-statistics/internal/common"
)

type StatisticsDatabase interface {
	GetPostStatistics(ctx context.Context, postId uint64) (*common.PostStatistics, error)
	GetTopKPosts(ctx context.Context, eventType string, k uint64) ([]common.PostStatistics, error)
	GetTopKUsers(ctx context.Context, eventType string, k uint64) ([]common.UserStatistics, error)
}

func NewDatabase(conn driver.Conn) StatisticsDatabase {
	return &StatisticsClickhouse{conn: conn}
}
