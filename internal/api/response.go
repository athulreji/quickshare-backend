package api

type GeneratePutUrlResponseBody struct {
	FileId       string `json:"fileId"`
	PresignedUrl string `json:"presignedUrl"`
}

type GenerateGetUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
	FileName     string `json:"fileName"`
}

type AddToDbResponseBody struct {
	Status string `json:"status"`
}
