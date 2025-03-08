package client

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrbelka12000/mock_server/pkg/config"
	"github.com/mrbelka12000/mock_server/pkg/server"
)

func Run() {
	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("app", "client")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	httpServer := server.Run(mux, cfg.ClientPort)
	waitCh := make(chan os.Signal)

	signal.Notify(waitCh, syscall.SIGINT, syscall.SIGTERM)

	log.With("port", cfg.ServerPort).Info("client started")
	select {
	case <-waitCh:
		log.Info("Interrupt signal received")
		httpServer.Close(context.Background())
	case <-httpServer.Wait():
		log.Info("Server exited")
	}

	log.Info("Shutting down")
}
