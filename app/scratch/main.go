package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	log := slog.Default()
	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.ErrorContext(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *slog.Logger) error {

	return nil
}
