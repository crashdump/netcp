package entity

type API struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content string `json:"content"`
}

type APIBlob struct {
	Filename string `json:"filename"`
	Content  string `json:"content"` // base64
}
