package models

type Code struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Input    string `json:"input"`
}
