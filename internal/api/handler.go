package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"quickshare-backend/internal/models"
	"quickshare-backend/internal/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "QuickShare API")
}

func GeneratePutUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	fileId, err := models.GenerateItemToDb()
	if err != nil {
		http.Error(w, "Unable to access db", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	url, err := services.GeneratePresignedPutUrl(fileId)
	if err != nil {
		http.Error(w, "Unable to access cloud storage", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := GeneratePutUrlResponseBody{
		FileId:       fileId,
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
	fileName, err := models.UserValidate(reqBody.FileId, reqBody.Password)

	if err != nil {
		http.Error(w, "Authentication Failed", http.StatusNotAcceptable)
		return
	}

	//fetch presigned get url and file name
	url, err := services.GeneratePresignedGetUrl(reqBody.FileId)
	if err != nil {
		http.Error(w, "Unable to access storage", http.StatusInternalServerError)
		return
	}
	response := GenerateGetUrlResponseBody{
		PresignedUrl: url,
		FileName:     fileName,
	}
	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)

}

func AddToDbHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var reqBody AddToDbRequestBody
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//Add FileName, Password to db
	err := models.UpdateDb(reqBody.FileId, reqBody.FileName, reqBody.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//return fileId
	w.Header().Set("Content-Type", "application/json")
	response := AddToDbResponseBody{
		Status: "Success",
	}
	responseJson, _ := json.Marshal(response)
	w.Write(responseJson)
}
