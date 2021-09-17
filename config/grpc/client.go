package grpc

import (
	microbackendfilm "micro_backend_film"
	pb_auth "micro_backend_film/services/auth/pb"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"micro_backend_film/services/biz/pb/pb_film"
	"micro_backend_film/services/bookmark/pb"
)

var (
	config         = microbackendfilm.Config()
	BizConn        = NewGrpcClient(config.Services["biz"].Port)
	BizGateClient  = pb_film.NewBizServiceClient(BizConn)
	BizCrawlClient = pb_crawl.NewBizServiceClient(BizConn)
	AuthConn       = NewGrpcClient(config.Services["auth"].Port)
	AuthClient     = pb_auth.NewAuthServiceClient(AuthConn)
	BMConn         = NewGrpcClient(config.Services["bookmark"].Port)
	BMClient       = pb.NewBookmarkServiceClient(BMConn)
)
