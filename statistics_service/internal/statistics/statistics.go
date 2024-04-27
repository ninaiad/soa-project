package statistics

import "soa-statistics/internal/database"

type StatisticsService struct {
	db database.Database
	// gRPC server stub
	// kafka producer
}

func NewStatisticsService(db database.Database) *StatisticsService {
	return &StatisticsService{db: db}
}
