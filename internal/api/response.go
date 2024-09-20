package api

type GeneratePostUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
}

type GenerateGetUrlResponseBody struct {
	PresignedUrl string `json:"presignedUrl"`
	FileName     string `json:"fileName"`
}

type AddToDbResponseBody struct {
	FileId string `json:"fileId"`
}
