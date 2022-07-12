package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"io/ioutil"
	"net/http"

	"demo/user/model"
	"demo/user/service"

	"demo/user/endpoint"
)

// Server
// 1. decode request      http.request -> model.request
func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("参数解析错误")
	}

	req := &model.CreateReq{}
	err = json.Unmarshal(bodyBytes, &req)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("参数解析错误")
	}
	fmt.Printf("%#v\n", req)

	err = r.Body.Close()
	if err != nil {
		return nil, err
	}
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
