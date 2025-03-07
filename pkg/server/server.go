package server

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	close      chan error
}

func Run(handler http.Handler, port int) *Server {
	server := &Server{
		httpServer: &http.Server{
			Handler: handler,
			Addr:    fmt.Sprintf(":%d", port),
		},
	}

	go func() {
		server.close <- server.httpServer.ListenAndServe()
	}()

	return server
}

func (s *Server) Close(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) Wait() <-chan error {
	return s.close
}
