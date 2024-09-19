package api

type GeneratePostUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
}

type GenerateGetUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
	FileName     string `json:"fileName"`
}

type AddPasswordResponseBody struct {
	FileId string `json:"fileId"`
}
