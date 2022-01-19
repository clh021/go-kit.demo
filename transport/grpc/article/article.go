package article

import (
	"context"
	endpoint "demo/endpoint/article"
	pb "demo/pb/article"
	"demo/service"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

// ArticleGrpcServer 1.实现了 pb.ArticleServiceServer 的所有方法，实现了”继承“;
// 2.提供了定义了 create 和 detail 两个 grpctransport.Handler。
type ArticleGrpcServer struct {
	createHandler grpctransport.Handler
	detailHandler grpctransport.Handler
}

// 通过 grpc 调用 Create 时，Create 只做数据传递, Create 内部又调用 createHandler，转交给 go-kit 处理

func (s *ArticleGrpcServer) Create (ctx context.Context, req *pb.CreateReq) (*pb.CreateResp, error) {
	_, rsp, err := s.createHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.CreateResp), err
}

func (s *ArticleGrpcServer) Detail (ctx context.Context, req *pb.DetailReq) (*pb.DetailResp, error) {
	_, rsp, err := s.detailHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.(*pb.DetailResp), err
}

// NewArticleGrpcServer 返回 proto 中定义的 article grpc server
func NewArticleGrpcServer(svc service.ArticleService, opts ...grpctransport.ServerOption) pb.ArticleServiceServer {

	createHandler := grpctransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		opts...,
	)

	articleGrpServer := new(ArticleGrpcServer)
	articleGrpServer.createHandler = createHandler

	return articleGrpServer
}
