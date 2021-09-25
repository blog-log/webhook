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
	AddApiPath = "/add"
)

type AddRequest struct {
	Repos []string `json:"repos"`
}

type Adder func(ctx context.Context, request *AddRequest) error

func (s *ProcessorService) Add(ctx context.Context, request *AddRequest) error {
	// construct request url
	url, err := url.Parse(fmt.Sprintf("%s%s", s.host, AddApiPath))
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
