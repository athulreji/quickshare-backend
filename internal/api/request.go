package api

type GeneratePostUrlRequestBody struct {
	FileName string `json:"fileName"`
}

type GenerateGetUrlRequestBody struct {
	FileId   string `json:"fileId"`
	Password string `json:"password"`
}

type AddPasswordRequestBody struct {
	FileName string `json:"fileName"`
	Password string `json:"password"`
}
