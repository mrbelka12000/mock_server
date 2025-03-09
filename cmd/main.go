package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/mrbelka12000/mock_server/internal/handler"
	"github.com/mrbelka12000/mock_server/internal/service"
	"github.com/mrbelka12000/mock_server/internal/storage"
	"github.com/mrbelka12000/mock_server/pkg/config"
	"github.com/mrbelka12000/mock_server/pkg/database"
	"github.com/mrbelka12000/mock_server/pkg/server"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		panic(err)
	}

	log := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("app", "sever")

	db, err := database.Connect(cfg)
	if err != nil {
		log.With("error", err).Error("connect to database")
		return
	}
	defer db.Close()

	store := storage.New(db)

	srv := service.New(store)
	dynHandler := handler.NewDynamicHandler(srv, handler.WithLogger(log.With("module", "handler")))
	httpServer := server.Run(dynHandler, cfg.ServerPort)

	waitCh := make(chan os.Signal)

	signal.Notify(waitCh, syscall.SIGINT, syscall.SIGTERM)

	log.With("port", cfg.ServerPort).Info("server started")
	select {
	case <-waitCh:
		log.Info("Interrupt signal received")
		httpServer.Close(context.Background())
	case <-httpServer.Wait():
		log.Info("Server exited")
	}

	log.Info("Shutting down")
}
