package github

import (
	"github.com/blog-log/webhook/internal/config"
	"github.com/blog-log/webhook/internal/processor"
)

type Handler struct {
	config    *config.Webhook
	processor processor.Processor
}

func NewHandler(cfg *config.Webhook) *Handler {
	return &Handler{
		config:    cfg,
		processor: processor.NewGithubProcessor(cfg),
	}
}
