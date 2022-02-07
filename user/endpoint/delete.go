package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"

	errors "demo/errors"
	"demo/user/model"
	"demo/user/service"
)

// make endpoint             service -> endpoint
func MakeDeleteEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*model.DeleteReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
