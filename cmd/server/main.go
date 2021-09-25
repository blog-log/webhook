package main

import (
	"context"
	"log"

	"github.com/blog-log/webhook/internal/config"
	serverv1 "github.com/blog-log/webhook/internal/server"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config", err)
	}

	server := serverv1.NewWebhookServer()

	server.Run(ctx, cfg)
}
