package main

import (
	"log"
	"micro_backend_film/config/cache"
	"micro_backend_film/config/db"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"micro_backend_film/services/biz/pb/pb_film"
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

	// DB
	db := db.ConnectPostgres()
	filmRepo := &repo.FilmRepo{
		DB: db,
	}

	// Redis
	red := cache.ConnectRedis()

	// Handler
	crawlHandler := &handler.CrawlHandler{RedisCache: red, FilmRepo: filmRepo}
	filmHandler := &handler.FilmHandler{RedisCache: red, FilmRepo: filmRepo}

	// GRPC
	grpcServer := grpc.NewServer()

	pb_crawl.RegisterBizServiceServer(grpcServer, crawlHandler)
	pb_film.RegisterBizServiceServer(grpcServer, filmHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 8081: %v", err)
	}
}
