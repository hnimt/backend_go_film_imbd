package grpc

import (
	pb_auth "micro_backend_film/services/auth/pb"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"micro_backend_film/services/biz/pb/pb_film"
	"micro_backend_film/services/bookmark/pb"
)

var (
	BizConn        = NewGrpcClient(":8081")
	BizGateClient  = pb_film.NewBizServiceClient(BizConn)
	BizCrawlClient = pb_crawl.NewBizServiceClient(BizConn)
	AuthConn       = NewGrpcClient(":8082")
	AuthClient     = pb_auth.NewAuthServiceClient(AuthConn)
	BMConn         = NewGrpcClient(":8083")
	BMClient       = pb.NewBookmarkServiceClient(BMConn)
)
