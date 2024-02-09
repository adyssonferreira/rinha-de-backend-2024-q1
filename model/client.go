package model

type Client struct {
	Id      string `sql:"id"`
	Nome    string `sql:"name"`
	Balance int    `sql:"balance"`
	Limit   int    `sql:"limit"`
}
