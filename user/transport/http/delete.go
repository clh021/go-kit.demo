package http

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"

	"demo/user/endpoint"
	"demo/user/model"
	"demo/user/service"
)

// Server
// 1. decode request      http.request -> model.request
func decodeDeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if err := FormCheckAccess(r); err != nil {
		return nil, err
	}
	r.ParseForm()
	req := &model.DeleteReq{}
	err := ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return req, nil
}

// 2. encode response      model.response -> http.response
func encodeDeleteResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// make handler
func MakeDeleteHandler(svc service.UserService) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeDeleteEndpoint(svc),
		decodeDeleteRequest,
		encodeDeleteResponse,
		ErrorServerOption(), // 自定义错误处理
	)
	return handler
}
