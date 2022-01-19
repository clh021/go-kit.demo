package grpc

import (
	"demo/router/grpcrouter"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

var opts = []grpc.ServerOption{
	grpc_middleware.WithUnaryServerChain(
		RecoveryInterceptor,
	),
}

var grpcServer = grpc.NewServer(opts...)

func Run(addr string, errc chan error) {

	// 注册grpcServer
	grpcrouter.RegisterRouter(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}

	errc <- grpcServer.Serve(lis)
}
