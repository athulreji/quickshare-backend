package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quickshare-backend/internal/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "QuickShare API")
}

func GeneratePostUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody GeneratePostUrlRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	url, err := services.GeneratePresignedPutUrl(reqBody.FileName)
	if err != nil {
		log.Print(err)
	}
	w.Header().Set("Content-Type", "application/json")
	response := GeneratePostUrlResponseBody{
		PresignedUrl: url,
	}
	responseJson, _ := json.Marshal(response)

	w.Write(responseJson)
}

func GenerateGetUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody GenerateGetUrlRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//validate fileid and password
	

	//fetch presigned get url and file name

}

func AddToDb(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody AddPasswordRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//Add FileName FileID to db


	//return fileId

}
