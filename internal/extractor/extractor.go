package extractor

import "context"

// extractor returns repos added, removed or an error
type Extractor interface {
	Extract(ctx context.Context, event interface{}) ([]string, []string, error)
}
