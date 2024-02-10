package model

import "time"

type Tansaction struct {
	Value       int       `json:"valor"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao"`
	CreateAt    time.Time `json:"realizada_em"`
}
