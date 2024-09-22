package main

import (
	"fmt"
	"net/http"
	"quickshare-backend/internal/api"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.HomeHandler)
	mux.HandleFunc("/generate_put_url", api.GeneratePutUrlHandler)
	mux.HandleFunc("/generate_get_url", api.GenerateGetUrlHandler)
	mux.HandleFunc("/add_to_db", api.AddToDbHandler)

	fmt.Println("Starting server on :5000")
	http.ListenAndServe(":5000", mux)
}
