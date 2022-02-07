package http

import (
	"demo/user/service"
	transport "demo/user/transport/http"
	"net"
	"net/http"
)

var mux = http.NewServeMux()

var httpServer = http.Server{Handler: mux}

func WrapMiddleware(handlerMap map[string]http.Handler, middlewares ...HttpHandlerMiddleware) map[string]http.Handler {
	m := map[string]http.Handler{}
	for k, h := range handlerMap {
		var newHandler http.Handler
		for _, m := range middlewares {
			newHandler = m(h)
		}
		m[k] = newHandler
	}
	return m
}

func RegisterRouter(mux *http.ServeMux)  {
	mux.Handle("/user/create", transport.MakeCreateHandler(service.NewUserService()))
	mux.Handle("/user/delete", transport.MakeDeleteHandler(service.NewUserService()))
}

// http run
func Run(addr string, errc chan error) {

	// 注册路由
	RegisterRouter(mux)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)
}

