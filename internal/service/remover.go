package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	RemoveApiPath = "/remove"
)

type RemoveRequest struct {
	Repos []string `json:"repos"`
}

type Remover func(ctx context.Context, request *RemoveRequest) error

func (s *ProcessorService) Remove(ctx context.Context, request *RemoveRequest) error {
	// construct request url
	url, err := url.Parse(fmt.Sprintf("%s%s", s.host, RemoveApiPath))
	if err != nil {
		return err
	}

	// marshal request
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// prepare request
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	// set headers
	req.Header.Set("Content-Type", "application/json")

	// invoke request
	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	// validate response status code
	if resp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("invalid status code %v for processor add request", resp.StatusCode)
	}

	return nil
}
