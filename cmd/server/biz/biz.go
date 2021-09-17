package main

import (
	"log"
	microbackendfilm "micro_backend_film"
	"micro_backend_film/common/repo"
	"micro_backend_film/config/cache"
	"micro_backend_film/config/db"
	"micro_backend_film/services/biz/handler"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"micro_backend_film/services/biz/pb/pb_film"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Config
	config := microbackendfilm.Config()

	lis, err := net.Listen(config.Services["biz"].Prot, config.Services["biz"].Port)
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	// DB
	db := db.ConnectPostgres(config)
	filmRepo := &repo.FilmRepo{
		DB: db,
	}

	// Redis
	red := cache.ConnectRedis(config)

	// Handler
	crawlHandler := &handler.CrawlHandler{RedisCache: red, FilmRepo: filmRepo}
	filmHandler := &handler.FilmHandler{RedisCache: red, FilmRepo: filmRepo}

	// GRPC
	grpcServer := grpc.NewServer()

	pb_crawl.RegisterBizServiceServer(grpcServer, crawlHandler)
	pb_film.RegisterBizServiceServer(grpcServer, filmHandler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server %v", err)
	}
}
