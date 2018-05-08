package undergroundF

import (
	"net/http"
)

type Router struct {
	URL     string
	handler func(http.ResponseWriter, *http.Request)
}

func AddRouter(url string, handler func(http.ResponseWriter, *http.Request)) Router {
	return Router{url, handler}
}
