package extractor

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/go-github/github"
)

type GithubExtractor struct{}

func NewGithubExtractor() *GithubExtractor {
	return &GithubExtractor{}
}

func (e *GithubExtractor) Extract(ctx context.Context, event interface{}) ([]string, []string, error) {
	switch event.(type) {
	case *github.PushEvent:
		return e.ExtractPushEvent(ctx, event)
	case *github.PullRequestEvent:
		return e.ExtractPullRequestEvent(ctx, event)
	case *github.InstallationEvent:
		return e.ExtractInstallationEvent(ctx, event)
	case *github.InstallationRepositoriesEvent:
		return e.ExtractInstallationRepositoriesEvent(ctx, event)
	default:
		return nil, nil, errors.New("unhandled webhook event type")
	}
}

func (e *GithubExtractor) ExtractPushEvent(ctx context.Context, event interface{}) ([]string, []string, error) {
	ghEvent, ok := event.(*github.PushEvent)
	if !ok {
		return nil, nil, errors.New("error casting github push event")
	}

	return []string{
		fmt.Sprintf("https://github.com/%s", *ghEvent.Repo.FullName),
	}, nil, nil
}

func (e *GithubExtractor) ExtractPullRequestEvent(ctx context.Context, event interface{}) ([]string, []string, error) {
	ghEvent, ok := event.(*github.PullRequestEvent)
	if !ok {
		return nil, nil, errors.New("error casting github pull request event")
	}

	return []string{
		fmt.Sprintf("https://github.com/%s", *ghEvent.Repo.FullName),
	}, nil, nil
}

func (e *GithubExtractor) ExtractInstallationEvent(ctx context.Context, event interface{}) ([]string, []string, error) {
	ghEvent, ok := event.(*github.InstallationEvent)
	if !ok {
		return nil, nil, errors.New("error casting github installation event")
	}

	var repos []string
	for _, repository := range ghEvent.Repositories {
		repos = append(repos, fmt.Sprintf("https://github.com/%s", *repository.FullName))
	}

	switch event.(map[string]interface{})["action"] {
	case "created":
		return repos, nil, nil
	case "deleted":
		return nil, repos, nil
	default:
		return nil, nil, errors.New("unhandled webhook installation action type")
	}
}

func (e *GithubExtractor) ExtractInstallationRepositoriesEvent(ctx context.Context, event interface{}) ([]string, []string, error) {
	ghEvent, ok := event.(*github.InstallationRepositoriesEvent)
	if !ok {
		return nil, nil, errors.New("error casting github installation repositories event")
	}

	var reposAdded []string
	for _, repository := range ghEvent.RepositoriesAdded {
		reposAdded = append(reposAdded, fmt.Sprintf("https://github.com/%s", *repository.FullName))
	}

	var reposRemoved []string
	for _, repository := range ghEvent.RepositoriesRemoved {
		reposRemoved = append(reposRemoved, fmt.Sprintf("https://github.com/%s", *repository.FullName))
	}

	return reposAdded, reposRemoved, nil
}
