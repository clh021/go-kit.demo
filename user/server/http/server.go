package http

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"demo/user/service"
	transport "demo/user/transport/http"
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
	r.Methods("POST").Path("/user/create").Handler(transport.MakeCreateHandler(service.NewUserService()))
	r.Methods("DELETE").Path("/user/delete/{id}").Handler(transport.MakeDeleteHandler(service.NewUserService()))
	// 健康检测接口
	r.Methods("GET", "POST").Path("/health").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-type", "application/json")
		writer.Write([]byte(`{"status": "ok"}`))
	})
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

