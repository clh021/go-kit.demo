package grpc

import (
	"context"
	"demo/user/endpoint"
	pb "demo/user/pb"
	"demo/user/service"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// UserGrpcServer 1.实现了 pb.UserServiceServer 的所有方法，实现了”继承“;
// 2.提供了定义了 create 和 detail 两个 grpctransport.Handler。
type UserGrpcServer struct {
	createHandler grpctransport.Handler
	deleteHandler grpctransport.Handler
}

// 通过 grpc 调用 Create 时，Create 只做数据传递, Create 内部又调用 createHandler，转交给 go-kit 处理

func (s *UserGrpcServer) Create (ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	_, rsp, err := s.createHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.CreateResp), err
}

func (s *UserGrpcServer) Delete (ctx context.Context, req *pb.DeleteReq) (*pb.DeleteResp, error) {
	_, rsp, err := s.deleteHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.DeleteResp), err
}

// NewUserGrpcServer 返回 proto 中定义的 user grpc server
func NewUserGrpcServer(svc service.UserService, opts ...grpctransport.ServerOption) pb.UserServiceServer {

	createHandler := grpctransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		opts...,
	)

	deleteHandler := grpctransport.NewServer(
		endpoint.MakeDeleteEndpoint(svc),
		decodeDeleteRequest,
		encodeDeleteResponse,
		opts...,
	)

	userGrpServer := new(UserGrpcServer)
	userGrpServer.createHandler = createHandler
	userGrpServer.deleteHandler = deleteHandler

	return userGrpServer
}