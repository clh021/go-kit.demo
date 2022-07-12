package grpc

import (
	pb "demo/user/pb"
	"demo/user/service"
	transport "demo/user/transport/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
)

var opts = []grpc.ServerOption{
	grpc_middleware.WithUnaryServerChain(
		RecoveryInterceptor,
	),
}

var grpcServer = grpc.NewServer(opts...)

func Run(addr string, errc chan error) {
	// 注册grpcServer
	pb.RegisterUserServiceServer(grpcServer, transport.NewUserGrpcServer(service.NewUserService()))
	// 注册健康检查Server
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	// prometheus metric 接口：grpc 实现 prometheus metric 仍然需要依赖 http 服务。
	// 实现 demo: https://github.com/grpc-ecosystem/go-grpc-prometheus/blob/master/examples/grpc-server-with-prometheus/server/server.go

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}

	errc <- grpcServer.Serve(lis)
}

