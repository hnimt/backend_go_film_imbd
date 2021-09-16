package main

import (
	"context"
	"log"
	"micro_backend_film/config/grpc"
	"micro_backend_film/services/crawler/handler"
	"time"

	"github.com/gocolly/colly"
)

func main() {

	colly := colly.NewCollector()
	handerCrawl := &handler.HandlerCrawl{
		Colly: colly,
	}

	msg := handerCrawl.CrawlFilm()

	ticker := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-ticker.C:
			response, err := grpc.BizCrawlClient.CrawlFilm(context.Background(), msg)
			if err != nil {
				log.Fatalf("Error when calling CrawlFilm: %s name", err)
			}
			log.Printf("Response from Server ", response)
		}
	}

}
