package article

import (
	"context"
	endpoint "demo/endpoint/article"
	"demo/params/article_param"
	"demo/service"
	transport "demo/transport/http"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

// Server
// 1. decode request      http.request -> model.request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	r.ParseForm()
	req := &article_param.CreateReq{}
	err := transport.ParseForm(r.Form, req)
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
func MakeCreateHandler(svc service.ArticleService) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}