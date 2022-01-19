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
func decodeDetailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	r.ParseForm()
	req := &article_param.DetailReq{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return req, nil
}

// 2. encode response      model.response -> http.response
func encodeDetailResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

// make handler
func MakeDetailHandler(svc service.ArticleService) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeDetailEndpoint(svc),
		decodeDetailRequest,
		encodeDetailResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}
