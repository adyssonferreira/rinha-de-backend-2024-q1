package model

type Tansaction struct {
	Value       int    `json:"valor"`
	Type        string `json:"tipo"`
	Description string `json:"descricao"`
	CreateAt    string `json:"realizada_em"`
}
