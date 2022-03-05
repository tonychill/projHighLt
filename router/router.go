package router

import "net/http"

type Router interface {
	GET(string, func(http.ResponseWriter, *http.Request))
	POST(string, func(http.ResponseWriter, *http.Request))
	SERVE(string)
}
