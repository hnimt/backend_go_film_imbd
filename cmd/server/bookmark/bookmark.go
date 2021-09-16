package main

import (
	"log"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/db"
	"micro_backend_film/services/bookmark/handler"
	"micro_backend_film/services/bookmark/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("Failed to listen on port 8083: %v", err)
	}

	// DB
	db := db.ConnectPostgres()
	bmRepo := &repo.BookmarkRepo{
		DB: db,
	}

	// Handler
	bmHandler := &handler.BookmarkHandler{BMRepo: bmRepo}

	// GRPC
	grpcServer := grpc.NewServer()

	pb.RegisterBookmarkServiceServer(grpcServer, bmHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8083: %v", err)
	}
}
