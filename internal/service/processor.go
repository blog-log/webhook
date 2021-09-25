package service

import (
	"crypto/tls"
	"net/http"
	"time"
)

// wrapper interface used for test mocking
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ProcessorService struct {
	host   string
	client HttpClient
}

func NewProcessorService(host string) *ProcessorService {
	// TODO - once get valid cert remove this
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return &ProcessorService{
		host: host,
		client: &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		},
	}
}
