package model

type Saldo struct {
	Total int    `json: total`
	Date  string `json: tipo`
	Limit int    `json: limite`
}
