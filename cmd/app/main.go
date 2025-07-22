package main

import (
	"context"
	"errors"

	"imageflow/internal/delivery"
	postgres "imageflow/internal/repository/postrges"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"imageflow/internal/config"
	"imageflow/internal/logger"
	"imageflow/internal/usecase"
)

// @title           ImageFlow API
// @version         1.0
// @description     API for submitting and checking image processing tasks.
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	logger.Init()
	slog.Info("Starting ImageFlow app")

	cfg := config.Load()
	repo, err := postgres.NewPostgresRepo("db", "5432", "postgres", "postgres", "imageflow")
	if err != nil {
		log.Fatal(err)
	}
	uc := usecase.NewImageUseCase(repo)
	handler := delivery.NewHandler(cfg, uc)

	go func() {
		if err := handler.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed: %v", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	slog.Warn("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := handler.Stop(ctx); err != nil {
		slog.Error("failed to shutdown gracefully", "err", err)
	} else {
		slog.Info("Server stopped cleanly")
	}
}
