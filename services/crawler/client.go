package crawler

import (
	"micro_backend_film/config/grpc"
	"micro_backend_film/services/biz/pb/pb_crawl"
)

var (
	Conn = grpc.NewGrpcClient(":8081")
	Client = pb_crawl.NewBizServiceClient(Conn)
)