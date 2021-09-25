package github

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/blog-log/webhook/internal/config"
	"github.com/google/go-github/github"
)

func (h *Handler) Validate(r *http.Request, cfg *config.Webhook) ([]byte, error) {
	if !cfg.GithubValidate {
		return h.validatePayloadOnly(r)
	}

	return github.ValidatePayload(r, []byte(h.config.GithubSecret))
}

// stolen from google github.ValidatePayload
func (h *Handler) validatePayloadOnly(r *http.Request) ([]byte, error) {
	var body []byte // Raw body that GitHub uses to calculate the signature.

	switch ct := r.Header.Get("Content-Type"); ct {
	case "application/json":
		var err error
		if body, err = ioutil.ReadAll(r.Body); err != nil {
			return nil, err
		}

		// If the content type is application/json,
		// the JSON payload is just the original body.
		return body, nil

	case "application/x-www-form-urlencoded":
		// payloadFormParam is the name of the form parameter that the JSON payload
		// will be in if a webhook has its content type set to application/x-www-form-urlencoded.
		const payloadFormParam = "payload"

		var err error
		if body, err = ioutil.ReadAll(r.Body); err != nil {
			return nil, err
		}

		// If the content type is application/x-www-form-urlencoded,
		// the JSON payload will be under the "payload" form param.
		form, err := url.ParseQuery(string(body))
		if err != nil {
			return nil, err
		}
		return []byte(form.Get(payloadFormParam)), nil

	default:
		return nil, fmt.Errorf("webhook request has unsupported Content-Type %q", ct)
	}
}
