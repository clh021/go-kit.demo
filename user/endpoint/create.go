package endpoint

import (
	"context"
	"demo/common/errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"

	"demo/user/model"
	"demo/user/service"
)

// make endpoint             service -> endpoint
func MakeCreateEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*model.CreateReq)
		fmt.Printf("%#v\n", *req)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
