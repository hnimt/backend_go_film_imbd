package gate

import (
	"micro_backend_film/config/grpc"
	pb_auth "micro_backend_film/services/auth/pb"
	"micro_backend_film/services/biz/pb/pb_film"
)

var (
	BizConn    = grpc.NewGrpcClient(":8081")
	BizClient  = pb_film.NewBizServiceClient(BizConn)
	AuthConn   = grpc.NewGrpcClient(":8082")
	AuthClient = pb_auth.NewAuthServiceClient(AuthConn)
)
