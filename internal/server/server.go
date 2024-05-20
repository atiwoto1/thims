package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/atiwoto1/thims/internal/config"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	cfc    *config.Config
	server *http.Server
	port   string
}

func NewServer(cfg *config.Config) *Server {
	port := cfg.GetEnvOrDefault("port", "8080")
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: routes(),
	}
	return &Server{
		cfc:    cfg,
		server: httpServer,
		port:   port,
	}
}

func (s *Server) Start(ctx context.Context) error {
	slog.Info(fmt.Sprintf("starting server on port %s", s.port))
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	return s.server.Shutdown(ctx)
}
