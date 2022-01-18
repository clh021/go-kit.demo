package order

import (
	"context"
	"demo/errors"
	"demo/params/article_param"
	service "demo/service"
	"github.com/go-kit/kit/endpoint"
)

// make endpoint             service -> endpoint
func MakeCreateEndpoint(svc service.ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*article_param.CreateReq)
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
