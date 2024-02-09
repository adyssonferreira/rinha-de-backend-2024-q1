package model

type Statement struct {
	Balance     Balance      `json:"saldo"`
	Tansactions []Tansaction `json:"ultimas_transacoes"`
}
