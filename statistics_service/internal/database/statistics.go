package database

import (
	"context"
	"fmt"

	"soa-statistics/internal/common"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type StatisticsClickhouse struct {
	conn driver.Conn
}

func (db *StatisticsClickhouse) GetPostStatistics(ctx context.Context, postId uint64) (*common.PostStatistics, error) {
	var dbPost common.PostStatistics
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

func (db *StatisticsClickhouse) GetTopKPosts(ctx context.Context, eventType string, k uint64) ([]common.PostStatistics, error) {
	var posts []common.PostStatistics
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

func (db *StatisticsClickhouse) GetTopKUsers(ctx context.Context, eventType string, k uint64) ([]common.UserStatistics, error) {
	var users []common.UserStatistics
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
