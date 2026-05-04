package models

import "time"

type Rutina struct {
    ID   int   `json:"id"`
    IDUsuario   int    `json:"id_usuario"`
    Nombre    string  `json:"nombre"`
    Descripcion   string  `json:"descripcion"`
    Objetivo    string  `json:"objetivo"`
    Dia   string  `json:"dia"` 
    Activo   bool  `json:"activo"`
    FechaModificacion time.Time `json:"fecha_modificacion"`
    FechaCreacion     time.Time `json:"fecha_creacion"`
}
