package model

type Extrato struct {
	Saldo      Saldo       `json: saldo`
	Transacoes []Transacao `json: ultimas_transacoes`
}
