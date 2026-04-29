package controllers

import (
	"api_rutinas_gym/config"
	"api_rutinas_gym/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)
// GetAllGruposMusculares obtener grupos musculares 
func GetAllGruposMusculares(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, nombre, descripcion, activo, fecha_modificacion, fecha_creacion FROM grupos_musculares WHERE 1=1"

	nombre := r.URL.Query().Get("nombre")
	if nombre != "" {
		query += " AND nombre LIKE '%" + nombre + "%'"
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.GrupoMuscular
	for rows.Next() {
		var g models.GrupoMuscular
		rows.Scan(&g.ID, &g.Nombre, &g.Descripcion, &g.Activo, &g.FechaModificacion, &g.FechaCreacion)
		list = append(list, g)
	}

	respondJSON(w, 200, list)
}

// GetGrupoMuscularByID obtener un grupo muscular por ID
func GetGrupoMuscularByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var g models.GrupoMuscular

	err := config.DB.QueryRow(
		"SELECT id, nombre, descripcion, activo, fecha_modificacion, fecha_creacion FROM grupos_musculares WHERE id = $1",
		id,
	).Scan(&g.ID, &g.Nombre, &g.Descripcion, &g.Activo, &g.FechaModificacion, &g.FechaCreacion)

	if err == sql.ErrNoRows {
		respondJSON(w, 404, map[string]string{"error": "Grupo muscular no encontrado"})
		return
	}
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, g)
}

// CreateGrupoMuscular crea un nuevo grupo muscular
func CreateGrupoMuscular(w http.ResponseWriter, r *http.Request) {
	var g models.GrupoMuscular
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	err := config.DB.QueryRow(
		"INSERT INTO grupos_musculares (nombre, descripcion, activo) VALUES ($1, $2, $3) RETURNING id, fecha_modificacion, fecha_creacion",
		g.Nombre, g.Descripcion, g.Activo,
	).Scan(&g.ID, &g.FechaModificacion, &g.FechaCreacion)

	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 201, g)
}

// UpdateGrupoMuscular actualiza un grupo muscular
func UpdateGrupoMuscular(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var g models.GrupoMuscular
	json.NewDecoder(r.Body).Decode(&g)

	_, err := config.DB.Exec(
		"UPDATE grupos_musculares SET nombre=$1, descripcion=$2, activo=$3 WHERE id=$4",
		g.Nombre, g.Descripcion, g.Activo, id,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Grupo muscular actualizado correctamente"})
}