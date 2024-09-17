package models

type GeneratePostUrlRequestBody struct {
	FileName string `json:"fileName"`
}

type GeneratePostUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
}
