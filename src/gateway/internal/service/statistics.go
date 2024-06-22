package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	stat_pb "gateway/internal/service/statistics_pb"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type EventType string

const (
	View EventType = "view"
	Like EventType = "like"
)

type KafkaConfig struct {
	KafkaAddr      string
	KafkaTopic     string
	KafkaGroupName string
}

func (s *Service) AddEvent(postId, authorId, actorId int64, eventType EventType) error {
	payload, err := json.Marshal(map[string]interface{}{
		"post_id":   postId,
		"author_id": authorId,
		"actor_id":  actorId,
		"event":     eventType,
		"timestamp": time.Now().Unix(),
	})
	if err != nil {
		return err
	}

	msg := &kafka.Message{
		Value: payload,
		TopicPartition: kafka.TopicPartition{
			Topic:     &s.kafkaCfg.KafkaTopic,
			Partition: kafka.PartitionAny,
		},
	}
	return s.kafkaProducer.Produce(msg, s.kafkaEventCh)
}

func (s *Service) GetPostStatistics(postId int64) (*stat_pb.PostStatistics, error) {
	return s.sClient.GetPostStatistics(context.Background(), &stat_pb.PostId{Id: postId})
}

func (s *Service) GetTopKPosts(event EventType, k uint64) (*stat_pb.TopPosts, error) {
	var eventType stat_pb.EventType
	if event == "like" {
		eventType = stat_pb.EventType_LIKE
	} else if event == "view" {
		eventType = stat_pb.EventType_VIEW
	} else {
		return nil, fmt.Errorf("unknown event_type value for get top k users")
	}

	return s.sClient.GetTopKPosts(context.Background(), &stat_pb.TopKRequest{Event: eventType, K: k})
}

func (s *Service) GetTopKUsers(event EventType, k uint64) (*stat_pb.TopUsers, error) {
	var eventType stat_pb.EventType
	if event == "like" {
		eventType = stat_pb.EventType_LIKE
	} else if event == "view" {
		eventType = stat_pb.EventType_VIEW
	} else {
		return nil, fmt.Errorf("unknown event_type value for get top k users")
	}

	return s.sClient.GetTopKUsers(context.Background(), &stat_pb.TopKRequest{Event: eventType, K: k})
}
