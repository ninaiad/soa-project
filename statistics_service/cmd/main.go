package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"soa-statistics/internal/common"
	"soa-statistics/internal/database"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	kafkaConfig := common.KafkaConfig{
		KafkaAddr:      os.Getenv("KAFKA_ADDR"),
		KafkaTopic:     viper.GetString("kafka.topic"),
		KafkaGroupName: viper.GetString("kafka.group_name"),
	}
	if kafkaConfig.KafkaAddr == "" {
		log.Fatalln("Missing KAFKA_ADDR")
	}

	log.Println("initialising the database next", viper.GetString("db.host"), viper.GetInt("db.port"))
	db, err := database.NewClickhouseDB(database.Config{
		Host:        viper.GetString("db.host"),
		Port:        viper.GetInt("db.port"),
		Username:    viper.GetString("db.username"),
		DBName:      viper.GetString("db.dbname"),
		Password:    os.Getenv("DB_PASSWORD"),
		KafkaCfg:    kafkaConfig,
		TableFormat: viper.GetString("db.table_format"),
	})
	if err != nil {
		log.Fatalf("error initialising clickhouse database: %s", err.Error())
	}
	log.Println("database is up and the migrations complete.")

	log.Println("initialising the HTTP server next")
	// for healthchecking
	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%s", viper.GetString("http_port")),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]any{"msg": "Hi! :)"})
		})}
	httpErrorCh := make(chan error, 1)
	go func() {
		httpErrorCh <- httpServer.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case err := <-httpErrorCh:
		log.Println("error from HTTP server: ", err.Error())
	}

	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	log.Println(httpServer.Shutdown(ctx).Error())
	log.Println(db.Close())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
