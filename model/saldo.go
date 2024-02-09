package model

type Balance struct {
	Total int    `json:"total"`
	Limit int    `json:"limite"`
	Date  string `json:"data_extrato"`
}
