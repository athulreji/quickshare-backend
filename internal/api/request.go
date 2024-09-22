package api

// type GeneratePutUrlRequestBody struct {
// 	FileName string `json:"fileName"`
// }

type GenerateGetUrlRequestBody struct {
	FileId   string `json:"fileId"`
	Password string `json:"password"`
}

type AddToDbRequestBody struct {
	FileId   string `json:"fileId"`
	FileName string `json:"fileName"`
	Password string `json:"password"`
}
