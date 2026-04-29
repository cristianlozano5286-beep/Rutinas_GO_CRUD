package models

import "time"

type GrupoMuscular struct {
	ID                int       `json:"id"`
	Nombre            string    `json:"nombre"`
	Descripcion       string    `json:"descripcion"`
	Activo            bool      `json:"activo"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
}
