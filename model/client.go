package model

type Client struct {
	Id      int    `sql:"id"`
	Nome    string `sql:"name"`
	Balance int    `sql:"balance"`
	Limit   int    `sql:"limit"`
}
