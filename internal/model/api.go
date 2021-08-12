package model

type APIError struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
