package main
import (
	"os"
	"net"
	"log"

	"soa/posts_service/internal/posts"
	pb "soa/posts"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // to be able to connect via grpcurl
)

func main() {
	// read server address from env
	addr := "0.0.0.0:"
	port := os.Getenv("POSTS_SERVER_PORT")
	if port == "" {
		addr = "0.0.0.0:51075"
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
	s := posts.NewPostsService()
	pb.RegisterPostsServerServer(grpcServer, s)

	// start the server
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("server failed")
	}
}