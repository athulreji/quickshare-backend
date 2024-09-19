package main

import (
	"fmt"
	"net/http"
	"quickshare-backend/internal/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.HomeHandler)
	mux.HandleFunc("/generate_post_url", api.GeneratePostUrlHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":5000", mux)
}
