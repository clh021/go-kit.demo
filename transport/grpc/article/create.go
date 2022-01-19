package article

import (
	"context"
	"demo/params/article_param"
	pb "demo/pb/article"
	"fmt"
)

// 1. decode request          pb -> model
func decodeCreateRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.CreateReq)
	if !ok {
		fmt.Println("grpc server decode request出错！")
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	// 过滤数据
	request := &article_param.CreateReq{
		Title: req.Title,
		Content: req.Content,
		CateId: req.CateId,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeCreateResponse(c context.Context, response interface{}) (interface{}, error) {
	fmt.Printf("%#v\n", response)
	resp, ok := response.(*article_param.CreateResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", response)
	}

	r := &pb.CreateResp{
		Id: resp.Id,
	}

	return r, nil
}
