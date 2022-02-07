package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"demo/user/model"
	"demo/user/service"

	"demo/user/endpoint"
)

// Server
// 1. decode request      http.request -> model.request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if err := FormCheckAccess(r); err != nil {
		return nil, err
	}
	r.ParseForm()
	req := &model.CreateReq{}
	err := ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return req, nil
}

// 2. encode response      model.response -> http.response
func encodeCreateResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// make handler
func MakeCreateHandler(svc service.UserService) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		ErrorServerOption(), // 自定义错误处理
	)
	return handler
}
