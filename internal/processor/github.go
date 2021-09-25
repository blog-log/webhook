package processor

import (
	"context"

	"github.com/blog-log/webhook/internal/config"
	"github.com/blog-log/webhook/internal/extractor"
	"github.com/blog-log/webhook/internal/service"
)

type GithubProcessor struct {
	extractor extractor.Extractor
	adder     service.Adder
	remover   service.Remover
}

func NewGithubProcessor(cfg *config.Webhook) *GithubProcessor {

	extractor := extractor.NewGithubExtractor()

	service := service.NewProcessorService(cfg.ProcessorHost)

	return &GithubProcessor{
		extractor: extractor,
		adder:     service.Add,
		remover:   service.Remove,
	}
}

func (p *GithubProcessor) Process(ctx context.Context, event interface{}) error {
	// extract out repo(s)
	reposAdded, reposRemoved, err := p.extractor.Extract(ctx, event)
	if err != nil {
		return err
	}

	// create or update repos
	if reposAdded != nil {
		if err := p.adder(ctx, &service.AddRequest{Repos: reposAdded}); err != nil {
			return err
		}
	}

	// remove repos
	if reposRemoved != nil {
		if err := p.remover(ctx, &service.RemoveRequest{Repos: reposRemoved}); err != nil {
			return err
		}
	}

	return nil
}
