package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"statistics/internal/db"
	"statistics/internal/kafka"
	pb "statistics/internal/pb"
	"statistics/internal/service"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // to be able to connect via grpcurl
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	kafkaConfig := kafka.Config{
		KafkaAddr:      os.Getenv("KAFKA_ADDR"),
		KafkaTopic:     viper.GetString("kafka.topic"),
		KafkaGroupName: viper.GetString("kafka.group_name"),
	}
	if kafkaConfig.KafkaAddr == "" {
		log.Fatalln("Missing KAFKA_ADDR")
	}

	log.Println("initialising the database next", viper.GetString("db.host"), viper.GetInt("db.port"))
	dbConn, err := db.NewClickhouseDB(db.Config{
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

	log.Println("initialising the gRPC server next")
	// read server address from env
	addr := "0.0.0.0:"
	port := os.Getenv("STATISTICS_SERVER_PORT")
	if port == "" {
		addr = "0.0.0.0:51706"
		log.Printf("Missing STATISTICS_SERVER_PORT, using default value: 51076")
	} else {
		addr += port
	}

	// create a TCP socket
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create & register the server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	s := service.NewStatisticsService(db.NewDatabase(dbConn))
	pb.RegisterStatisticsServiceServer(grpcServer, s)

	// start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("server failed")
	}

	log.Println("the gRPC is set up")

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
	log.Println(dbConn.Close())
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
