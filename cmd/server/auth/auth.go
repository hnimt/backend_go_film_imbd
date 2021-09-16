package main

import (
	"log"
	"micro_backend_film/config/cache"
	"micro_backend_film/config/db"
	"micro_backend_film/services/auth/handler"
	"micro_backend_film/services/auth/pb"
	"micro_backend_film/common/repo"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen on port 8082: %v", err)
	}

	// DB
	db := db.ConnectPostgres()
	userRepo := &repo.UserRepo{
		DB: db,
	}

	// Redis
	red := cache.ConnectRedis()

	// Handler
	authHander := &handler.AuthHandler{
		RdCache:  red,
		UserRepo: userRepo,
	}

	// GRPC
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authHander)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8082: %v", err)
	}
}
