package app

import (
	"context"
	"example/hello/internal/config"
	"example/hello/internal/handler"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) error {

	cfg := config.NewConfig()

	r := chi.NewRouter()

	handler.RegisterRoutes(r, handler.Dependencies{
		AssetsFS: http.Dir(cfg.AssetsDir),
	})

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
