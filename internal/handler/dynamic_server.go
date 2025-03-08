package handler

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/mrbelka12000/mock_server/internal"
	"github.com/mrbelka12000/mock_server/internal/service"
)

const (
	templatePath = "html_templates/"
)

type (
	DynamicRouter struct {
		srv       *service.Service
		mockPaths map[string]map[string]struct{}
		updateCh  chan routeInfo
		mu        sync.RWMutex
		log       *slog.Logger
	}

	routeInfo struct {
		serviceName string
		route       string
	}
)

func NewDynamicHandler(srv *service.Service, opts ...opt) *DynamicRouter {
	dr := &DynamicRouter{
		srv:       srv,
		mockPaths: make(map[string]map[string]struct{}),
		updateCh:  make(chan routeInfo),
		log:       slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}

	for _, opt := range opts {
		opt(dr)
	}

	dr.start()

	return dr
}

func (h *DynamicRouter) start() {
	services, err := h.srv.ListServices(context.Background())
	if err != nil {
		panic(err)
	}

	for _, service := range services {
		h.processer(service)
	}
}

func (h *DynamicRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")

	switch r.URL.Path {
	case "/service":
		h.HandleService(w, r)
	case "/handler":
		h.HandleHandlers(w, r)
	case "/":
		h.Index(w, r)
	default:
		paths := strings.Split(strings.TrimLeft(r.URL.Path, "/"), "/")
		if len(paths) != 3 {
			http.NotFound(w, r)
			return
		}

		serviceName := paths[1]

		routes, ok := h.mockPaths[serviceName]
		if !ok {
			http.NotFound(w, r)
			return
		}

		route := paths[2]
		_, ok = routes[route]
		if !ok {
			http.NotFound(w, r)
			return
		}
		log := h.log.With("service", serviceName, "route", route)

		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.With("error", err).Error("failed to read body")
			return
		}
		defer r.Body.Close()

		respBody, respHeader, err := h.srv.HandleRoute(r.Context(), serviceName, route, reqBody, r.Header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.With("error", err).Error("failed to handle route")
			return
		}

		for k, v := range respHeader {
			w.Header()[k] = v
		}

		w.WriteHeader(200)
		w.Write(respBody)
	}
}

func (h *DynamicRouter) Index(w http.ResponseWriter, r *http.Request) {
	
}

// processer adds new paths to router
func (h *DynamicRouter) processer(service internal.Service) {
	h.mu.Lock()
	defer h.mu.Unlock()

	_, ok := h.mockPaths[service.Name]
	if !ok {
		h.mockPaths[service.Name] = make(map[string]struct{})
	}

	for _, route := range service.Handlers {
		h.mockPaths[service.Name][route.Route] = struct{}{}
	}
}
