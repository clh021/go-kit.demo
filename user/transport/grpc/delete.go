package grpc

import (
	"context"
	"fmt"

	"demo/user/model"
	pb "demo/user/pb"
)

// Server
// 1. decode request          pb -> model
func decodeDeleteRequest(c context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pb.DeleteReq)
	if !ok {
		return nil, fmt.Errorf("grpc server decode request出错！")
	}
	request := &model.DeleteReq{
		Name: req.Name,
		Id:   req.Id,
	}
	return request, nil
}

// 2. encode response           model -> pb
func encodeDeleteResponse(c context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*model.DeleteResp)
	if !ok {
		return nil, fmt.Errorf("grpc server encode response出错！")
	}
	r := &pb.DeleteResp{
		Code: resp.Code,
		Msg:  resp.Msg,
		Data: &pb.DeleteRespData{
			Result: resp.Data.Result,
		},
	}
	return r, nil
}