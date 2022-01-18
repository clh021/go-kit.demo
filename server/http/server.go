package http

import (
	"demo/router/httprouter"
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

// http run
func Run(addr string, errc chan error) {

	// 注册路由
	httprouter.RegisterRouter(mux)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)
}
