package controllers

import (
	"api_rutinas_gym/config"
	"api_rutinas_gym/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
// GetAllEjercicios de los ejercicios 
func GetAllEjercicios(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, grupo_muscular_id, nombre, descripcion_corta, descripcion_larga,
	           posicion_inicial, ejecucion, consejos, nivel, activo, fecha_modificacion, fecha_creacion
	           FROM ejercicios WHERE 1=1`

	nombre := r.URL.Query().Get("nombre")
	nivel := r.URL.Query().Get("nivel")
	grupoID := r.URL.Query().Get("grupo_muscular_id")

	if nombre != "" {
		query += " AND nombre LIKE '%" + nombre + "%'"
	}
	if nivel != "" {
		query += " AND nivel = '" + nivel + "'"
	}
	if grupoID != "" {
		query += " AND grupo_muscular_id = " + grupoID
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Ejercicio
	for rows.Next() {
		var e models.Ejercicio
		rows.Scan(
			&e.ID, &e.GrupoMuscularID, &e.Nombre, &e.DescripcionCorta, &e.DescripcionLarga,
			&e.PosicionInicial, &e.Ejecucion, &e.Consejos, &e.Nivel, &e.Activo,
			&e.FechaModificacion, &e.FechaCreacion,
		)
		list = append(list, e)
	}

	respondJSON(w, 200, list)
}