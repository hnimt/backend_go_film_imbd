package main

import (
	"log"
	"micro_backend_film/config"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/db"
	"micro_backend_film/services/bookmark/handler"
	"micro_backend_film/services/bookmark/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Config
	config := config.Config()

	lis, err := net.Listen(config.Services["bookmark"].Prot, config.Services["bookmark"].Port)
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", err, config.Services["bookmark"].Port)
	}

	// DB
	db := db.Connect(config)
	bmRepo := &repo.BookmarkRepo{
		DB: db,
	}

	// Handler
	bmHandler := &handler.BookmarkHandler{BMRepo: bmRepo}

	// GRPC
	grpcServer := grpc.NewServer()

	pb.RegisterBookmarkServiceServer(grpcServer, bmHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %v: %v", err, config.Services["bookmark"].Port)
	}
}
