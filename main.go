package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/henrywhitakercommify/cfimport/cmd/root"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cmd := root.New()
	cmd.SetContext(ctx)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
