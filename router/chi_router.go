package router

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi/v5"
)

type ChiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &ChiRouter{}
}

func (*ChiRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Get(uri, f)
}
func (*ChiRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	chiDispatcher.Post(uri, f)

}
func (*ChiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v.", port)
	http.ListenAndServe(port, chiDispatcher)
}
