package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := github.ValidatePayload(r, []byte("1234567"))
	if err != nil {
		log.Printf("error validating request body: err=%s\n", err)
		fmt.Fprintf(w, "Failed to validate request body")
		return
	}
	defer r.Body.Close()

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/payload", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
