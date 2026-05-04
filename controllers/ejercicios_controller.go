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
	           FROM rutinas.ejercicios WHERE 1=1`

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
// GetEjercicioByID obtiene un ejercicio por ID
func GetEjercicioByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var e models.Ejercicio

	err := config.DB.QueryRow(
		`SELECT id, grupo_muscular_id, nombre, descripcion_corta, descripcion_larga,
		 posicion_inicial, ejecucion, consejos, nivel, activo, fecha_modificacion, fecha_creacion
		 FROM rutinas.ejercicios WHERE id = $1`, id,
	).Scan(
		&e.ID, &e.GrupoMuscularID, &e.Nombre, &e.DescripcionCorta, &e.DescripcionLarga,
		&e.PosicionInicial, &e.Ejecucion, &e.Consejos, &e.Nivel, &e.Activo,
		&e.FechaModificacion, &e.FechaCreacion,
	)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Ejercicio no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, e)
}
// CreateEjercicio crea un nuevo ejercicio
func CreateEjercicio(w http.ResponseWriter, r *http.Request) {
	var e models.Ejercicio
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		`INSERT INTO rutinas.ejercicios (grupo_muscular_id, nombre, descripcion_corta, descripcion_larga,
    posicion_inicial, ejecucion, consejos, nivel, activo)
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
    RETURNING id, fecha_modificacion, fecha_creacion`,
		e.GrupoMuscularID, e.Nombre, e.DescripcionCorta, e.DescripcionLarga,
		e.PosicionInicial, e.Ejecucion, e.Consejos, e.Nivel, e.Activo,
	).Scan(&e.ID, &e.FechaModificacion, &e.FechaCreacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, e)
}
// UpdateEjercicio actualiza un ejercicio
func UpdateEjercicio(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var e models.Ejercicio
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err := config.DB.Exec(
		`UPDATE rutinas.ejercicios SET grupo_muscular_id=$1, nombre=$2, descripcion_corta=$3,
		 descripcion_larga=$4, posicion_inicial=$5, ejecucion=$6, consejos=$7, nivel=$8, activo=$9
		 WHERE id=$10`,
		e.GrupoMuscularID, e.Nombre, e.DescripcionCorta, e.DescripcionLarga,
		e.PosicionInicial, e.Ejecucion, e.Consejos, e.Nivel, e.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Ejercicio actualizado correctamente"})
}

// DeleteEjercicio elimina un ejercicio
func DeleteEjercicio(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := config.DB.Exec("DELETE FROM rutinas.ejercicios WHERE id=$1", id)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Ejercicio eliminado correctamente"})
}