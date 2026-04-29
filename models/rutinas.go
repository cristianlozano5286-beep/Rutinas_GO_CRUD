package models

import "time"

type Rutina struct {
	IDUsuario         int       `json:"id_usuario"`
	Nombre            string    `json:"nombre"`
	Descripcion       string    `json:"descripcion"`
	Objetivo          string    `json:"objetivo"`
	Activo            bool      `json:"activo"`
	FechaModificacion time.Time `json:"fecha_modificacion"`
	FechaCreacion     time.Time `json:"fecha_creacion"`
}
