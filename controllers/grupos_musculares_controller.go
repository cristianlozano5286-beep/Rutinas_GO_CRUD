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