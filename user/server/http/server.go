package http

import (
	"net"
	"net/http"

	"demo/user/service"
	transport "demo/user/transport/http"

	"github.com/gorilla/mux"
)

//var mux = http.NewServeMux()
//
//var httpServer = http.Server{Handler: mux}
var httpServer = http.Server{}

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

func RegisterRouter(r *mux.Router)  {
	r.Handle("/user/create", transport.MakeCreateHandler(service.NewUserService())).Methods("POST")
	r.Handle("/user/delete/{id}", transport.MakeDeleteHandler(service.NewUserService())).Methods("DELETE")
}

// http run
func Run(addr string, errc chan error) {
	// 注册路由
	r := mux.NewRouter()
	RegisterRouter(r)
	httpServer.Handler = r

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		errc <- err
		return
	}
	errc <- httpServer.Serve(lis)
}

