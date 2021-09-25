package github

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/go-github/github"
)

func (h *Handler) Injest(w http.ResponseWriter, r *http.Request) {
	payload, err := h.Validate(r, h.config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	event, err := github.ParseWebHook(github.WebHookType(r), payload)
	if err != nil {
		http.Error(w, errors.New("failed to parse github webhook payload").Error(), http.StatusBadRequest)
		return
	}

	if err := h.processor.Process(r.Context(), event); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// everything went well
	w.WriteHeader(200)
	fmt.Fprint(w, "success")
}
