package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (*chiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)
}
func (*chiRouter) SERVE(port string, f func(w http.ResponseWriter, r *http.Request)) {
	fmt.printf("Chi HTTP Server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
