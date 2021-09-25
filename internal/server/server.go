package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/blog-log/webhook/internal/config"
	"github.com/blog-log/webhook/internal/handler/github"
	"github.com/gorilla/mux"
)

type Server interface {
	Run(ctx context.Context) error
}

type WebhookServer struct{}

func NewWebhookServer() *WebhookServer {
	return &WebhookServer{}
}

func (s *WebhookServer) Run(ctx context.Context, cfg *config.Webhook) error {

	rootRouter := mux.NewRouter().StrictSlash(true)

	// setup github sub router
	ghRouter := rootRouter.PathPrefix("/github").Subrouter()

	// setup GithubHandler
	handler := github.NewHandler(cfg)

	// define github route(s)
	ghRouter.HandleFunc("/injest", handler.Injest).Methods("POST")

	fmt.Println("Server started at port 8080")
	if err := http.ListenAndServe(":8080", rootRouter); err != nil {
		return err
	}

	return nil
}
