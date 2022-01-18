package httprouter

import (
	svc "demo/service"
	transport "demo/transport/http/article"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux)  {
	mux.Handle("/article/create", transport.MakeCreateHandler(svc.NewArticleService()))
}
