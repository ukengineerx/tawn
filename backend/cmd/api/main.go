package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"tawn/backend/internal/config"
	"tawn/backend/internal/httpapi"
	"tawn/backend/internal/notes"
	"tawn/backend/internal/platform/postgres"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.Load()
	if cfg.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	store := openStore(cfg.DatabaseURL)

	slog.Info("api listening", slog.String("port", cfg.Port), slog.String("env", cfg.Env))
	if err := httpapi.New(store).Run(":" + cfg.Port); err != nil {
		slog.Error("server stopped", slog.Any("error", err))
		os.Exit(1)
	}
}

func openStore(databaseURL string) notes.Store {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := postgres.Open(ctx, databaseURL)
	if err != nil {
		slog.Warn("postgres unavailable; /health stays up and /v1/notes returns 503", slog.Any("error", err))
		return nil
	}
	if err := postgres.Migrate(ctx, pool); err != nil {
		slog.Error("migration failed", slog.Any("error", err))
		os.Exit(1)
	}
	slog.Info("postgres ready")
	return postgres.NewStore(pool)
}
