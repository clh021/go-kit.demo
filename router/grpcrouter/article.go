package grpcrouter

import (
	pb "demo/pb/article"
	"demo/service"
	transport "demo/transport/grpc/article"
	"google.golang.org/grpc"
)

func RegisterRouter(grpcServer *grpc.Server) {
	pb.RegisterArticleServiceServer(grpcServer, transport.NewArticleGrpcServer(service.NewArticleService()))
}
