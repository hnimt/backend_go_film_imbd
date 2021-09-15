package grpc

import (
	"log"

	"google.golang.org/grpc"
)

func NewGrpcClient(port string) *grpc.ClientConn {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect :%s", err)
	}

	return conn
}
