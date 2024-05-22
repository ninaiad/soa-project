package database

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Database interface {
}

type StatisticsDatabase struct {
	conn driver.Conn
}

func NewDatabase(conn driver.Conn) Database {
	return &StatisticsDatabase{conn: conn}
}
