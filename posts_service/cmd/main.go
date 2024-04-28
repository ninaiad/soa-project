package main

import (
	"log"
	"net"
	"os"

	"soa-posts/internal/database"
	"soa-posts/internal/posts"
	pb "soa-posts/internal/proto"

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

	// read server address from env
	addr := "0.0.0.0:"
	port := os.Getenv("POSTS_SERVER_PORT")
	if port == "" {
		addr = "0.0.0.0:51705"
		log.Printf("Missing POSTS_SERVER_PORT, using default value: 51075")
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
	s := posts.NewPostsService(database.NewDatabase(db))
	pb.RegisterPostsServerServer(grpcServer, s)

	// start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("server failed")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
