package grpc

import (
	"micro_backend_film/config"
	pb_auth "micro_backend_film/services/auth/pb"
	"micro_backend_film/services/biz/pb/pb_crawl"
	"micro_backend_film/services/biz/pb/pb_film"
	"micro_backend_film/services/bookmark/pb"
)

var (
	tomlConfig     = config.Config()
	BizConn        = NewGrpcClient(tomlConfig.Services["biz"].Port)
	BizGateClient  = pb_film.NewBizServiceClient(BizConn)
	BizCrawlClient = pb_crawl.NewBizServiceClient(BizConn)
	AuthConn       = NewGrpcClient(tomlConfig.Services["auth"].Port)
	AuthClient     = pb_auth.NewAuthServiceClient(AuthConn)
	BMConn         = NewGrpcClient(tomlConfig.Services["bookmark"].Port)
	BMClient       = pb.NewBookmarkServiceClient(BMConn)
)
