package main

import (
	"log"
	"micro_backend_film/config/cache"
	"micro_backend_film/config/db"
	"micro_backend_film/services/biz/grpc/pb"
	"micro_backend_film/services/biz/handler"
	"micro_backend_film/services/biz/repo"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen on port 8081: %v", err)
	}

	db := db.ConnectPostgres()
	filmRepo := &repo.FilmRepo{
		DB: db,
	}

	red := cache.ConnectRedis()

	crawlHandler := &handler.CrawlHandler{RedisCache: red, FilmRepo: filmRepo}

	grpcServer := grpc.NewServer()

	pb.RegisterCrawlServiceServer(grpcServer, crawlHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8082: %v", err)
	}
}
