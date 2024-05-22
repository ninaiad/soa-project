package database

import (
	"context"
	"fmt"
	"soa-statistics/internal/common"

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
	KafkaCfg    common.KafkaConfig
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

	if err = conn.Exec(context.Background(),
		fmt.Sprintf(createTableQueue, cfg.KafkaCfg.KafkaAddr, cfg.KafkaCfg.KafkaTopic, cfg.KafkaCfg.KafkaGroupName, cfg.TableFormat)); err != nil {
		return nil, err
	}
	if err = conn.Exec(context.Background(), createTableEvents); err != nil {
		return nil, err
	}
	return conn, conn.Exec(context.Background(), createView)
}
