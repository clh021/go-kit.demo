package grpc

import (
	"context"
	"fmt"

	"demo/user/model"
	pb "demo/user/pb"
)

// Server
// 1. decode request          pb -> model
func decodeCreateRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.CreateReq)
	if !ok {
		fmt.Println("grpc server decode request出错！")
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.CreateReq{
		Name: req.Name,
		Age:  req.Age,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeCreateResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.CreateResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response error (%T)", response)
	}
	r := &pb.CreateResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.CreateRespData{
			Id:   resp.Data.Id,
			Age:  resp.Data.Age,
			Name: resp.Data.Name,
		},
	}
	return r, nil
}
