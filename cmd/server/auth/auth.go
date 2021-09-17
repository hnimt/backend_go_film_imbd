package main

import (
	"log"
	microbackendfilm "micro_backend_film"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/cache"
	"micro_backend_film/config/db"
	"micro_backend_film/services/auth/handler"
	"micro_backend_film/services/auth/pb"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Config
	config := microbackendfilm.Config()

	lis, err := net.Listen(config.Services["auth"].Prot, config.Services["auth"].Port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// DB
	db := db.ConnectPostgres(config)
	userRepo := &repo.UserRepo{
		DB: db,
	}

	// Redis
	red := cache.ConnectRedis(config)

	// Handler
	authHander := &handler.AuthHandler{
		RdCache:  red,
		UserRepo: userRepo,
	}

	// GRPC
	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, authHander)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server %v", err)
	}
}
