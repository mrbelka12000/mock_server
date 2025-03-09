package handler

import "net/http"

const (
	indexHTML   = "./web/public/index.html"
	serviceHTML = "./web/public/service.html"
)

func (h *DynamicRouter) HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, indexHTML)
}

func (h *DynamicRouter) HandleServicePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, serviceHTML)
}
