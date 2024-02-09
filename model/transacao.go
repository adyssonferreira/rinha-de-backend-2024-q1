package model

type Transacao struct {
	Valor    int    `json: valor`
	Tipo     string `json: tipo`
	Desc     int    `json: descricao`
	CreateAt string `json: realizada_em`
}
