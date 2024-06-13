package service

import (
	"encoding/json"
	"time"

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

type KafkaService struct {
	producer *kafka.Producer
	cfg      KafkaConfig
	eventCh  chan kafka.Event
}

func CreateKafkaService(p *kafka.Producer, cfg KafkaConfig, ch chan kafka.Event) *KafkaService {
	return &KafkaService{producer: p, cfg: cfg, eventCh: ch}
}

func (k *KafkaService) AddEvent(postId int64, authorId int64, eventType EventType) error {
	payload, err := json.Marshal(map[string]interface{}{
		"post":      postId,
		"author":    authorId,
		"event":     eventType,
		"timestamp": time.Now().Unix(),
	})
	if err != nil {
		return err
	}
	msg := &kafka.Message{
		Value:          payload,
		TopicPartition: kafka.TopicPartition{Topic: &k.cfg.KafkaTopic, Partition: kafka.PartitionAny},
	}
	return k.producer.Produce(msg, k.eventCh)
}
