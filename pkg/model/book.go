package model

import "time"

type Book struct {
	Titulo   string `json:"titulo"`
	Autor    string `json:"autor"`
	Año      int64  `json:"año"`
	CreadoEn time.Time `json:"creado_en"`
}
