package main

import (
	"context"
	"fmt"
	"github.com/atiwoto1/thims/internal/config"
	"github.com/atiwoto1/thims/internal/server"
	"io"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	if err := run(context.Background(), os.Getenv, os.Stderr, os.Stdout); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "\b", err)
		os.Exit(1)
	}
}

func run(
	ctx context.Context,
	getEnv func(string) string,
	stderr, stdout io.Writer,
) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	logger := slog.New(slog.NewJSONHandler(stdout, nil))
	slog.SetDefault(logger)

	conf := config.NewConfig(getEnv)
	srv := server.NewServer(conf)
	go func() {
		if err := srv.Start(ctx); err != nil {
			_, _ = fmt.Fprintln(stderr, "\n", err)
			cancel()
		}
	}()

	<-ctx.Done()

	slog.Info("shutting down...")
	return srv.Stop(context.Background())
}
