package main

import (
	"fmt"
	"net/http"
	"quickshare-backend/internal/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/generate_post_url", handlers.GeneratePostUrlHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}
