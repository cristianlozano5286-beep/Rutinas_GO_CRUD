package models

import "time"

type Ejercicio struct {
	ID   int       `json:"id"`
	GrupoMuscularID   int `json:"grupo_muscular_id"`
	Nombre   string    `json:"nombre"`
	DescripcionCorta string `json:"descripcion_corta"`
	DescripcionLarga string `json:"descripcion_larga"`
	PosicionInicial  string `json:"posicion_inicial"`
	Ejecucion  string `json:"ejecucion"`
	Consejos string  `json:"consejos"`
	Nivel  string  `json:"nivel"`
	Activo  bool  `json:"activo"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
	FechaCreacion   time.Time `json:"fecha_creacion"`
}