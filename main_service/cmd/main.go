package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"soa-main/internal/database"
	"soa-main/internal/handler"
	"soa-main/internal/service"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	log.Println("opening db next")

	db, err := database.NewPostgresDB(database.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	// initialise Kafka
	kafkaConfig := service.KafkaConfig{
		KafkaAddr:      os.Getenv("KAFKA_ADDR"),
		KafkaTopic:     viper.GetString("kafka.topic"),
		KafkaGroupName: viper.GetString("kafka.group_name"),
	}
	if kafkaConfig.KafkaAddr == "" {
		log.Fatalln("Missing KAFKA_ADDR")
		return
	}

	// connect to Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaConfig.KafkaAddr})
	if err != nil {
		log.Fatalf("error initialising kafka producer: %s", err.Error())
	}
	defer p.Close()

	// initialise grpc client
	postsC, err := grpc.Dial(os.Getenv("POSTS_SERVER_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer postsC.Close()

	repos := database.NewDatabase(db)

	kafkaEventCh := make(chan kafka.Event)
	services := service.NewService(repos, postsC, p, kafkaConfig, kafkaEventCh)

	handlers := handler.NewHandler(services)

	srv := new(Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.SetupRouter()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-quit:
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				log.Fatalf("error occured on server shutting down: %s", err.Error())
			}

			if err := db.Close(); err != nil {
				log.Fatalf("error occured on db connection close: %s", err.Error())
			}

			return
		case e := <-kafkaEventCh:
			m := e.(*kafka.Message)
			if m.TopicPartition.Error != nil {
				log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
			} else {
				log.Printf("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
			}
		}
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
