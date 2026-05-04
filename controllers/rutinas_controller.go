package controllers

import (
	"api_rutinas_gym/config"
	"api_rutinas_gym/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
// GetAllRutinas obtener todas las rutinas 
func GetAllRutinas(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, id_usuario, nombre, descripcion, objetivo, dia, activo, fecha_modificacion, fecha_creacion 
	          FROM rutinas.rutinas WHERE 1=1`

	nombre := r.URL.Query().Get("nombre")
	objetivo := r.URL.Query().Get("objetivo")
	usuarioID := r.URL.Query().Get("id_usuario")

	if nombre != "" {
		query += " AND nombre LIKE '%" + nombre + "%'"
	}
	if objetivo != "" {
		query += " AND objetivo LIKE '%" + objetivo + "%'"
	}
	if usuarioID != "" {
		query += " AND id_usuario = " + usuarioID
	}

	rows, err := config.DB.Query(query)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

var list []models.Rutina
    for rows.Next() {
        var ru models.Rutina
        err := rows.Scan(&ru.ID, &ru.IDUsuario, &ru.Nombre, &ru.Descripcion, &ru.Objetivo, &ru.Dia, &ru.Activo, &ru.FechaModificacion, &ru.FechaCreacion)
        if err != nil {
            log.Println("Error al escanear fila:", err)
            continue
        }
        list = append(list, ru)
    }
	respondJSON(w, 200, list)
}

// GetRutinaByUsuario obtiene rutinas por ID de usuario
func GetRutinaByUsuario(w http.ResponseWriter, r *http.Request) {
	idUsuario := mux.Vars(r)["id_usuario"]

	rows, err := config.DB.Query(
		`SELECT id_usuario, nombre, descripcion, objetivo, dia, activo, fecha_modificacion, fecha_creacion
          FROM rutinas.rutinas WHERE id_usuario = $1`, idUsuario,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}
	defer rows.Close()

	var list []models.Rutina
	for rows.Next() {
		var ru models.Rutina
		rows.Scan(&ru.IDUsuario, &ru.Nombre, &ru.Descripcion, &ru.Objetivo, &ru.Activo, &ru.FechaModificacion, &ru.FechaCreacion)
		list = append(list, ru)
	}

	if len(list) == 0 {
		respondJSON(w, 404, map[string]string{"error": "No se encontraron rutinas para este usuario"})
		return
	}

	respondJSON(w, 200, list)
}
// CreateRutina crea una nueva rutina
func CreateRutina(w http.ResponseWriter, r *http.Request) {
	var ru models.Rutina
	if err := json.NewDecoder(r.Body).Decode(&ru); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}
// Definimos el query
    query := `INSERT INTO rutinas.rutinas (id_usuario, nombre, descripcion, objetivo, dia, activo)
              VALUES ($1, $2, $3, $4, $5, $6)
              RETURNING id, id_usuario, nombre, descripcion, objetivo, dia, activo, fecha_modificacion, fecha_creacion`

    // Ejecutamos y escaneamos los resultados en el objeto 'ru'
   // 1. Ejecutamos el QueryRow pasándole el query y los datos de la rutina (ru)
// 2. PEGADO al paréntesis de cierre, llamamos al .Scan()
err := config.DB.QueryRow(query, 
    ru.IDUsuario, 
    ru.Nombre, 
    ru.Descripcion, 
    ru.Objetivo, 
    ru.Dia, 
    ru.Activo,
).Scan(
    &ru.ID, 
    &ru.IDUsuario, 
    &ru.Nombre, 
    &ru.Descripcion, 
    &ru.Objetivo, 
    &ru.Dia, 
    &ru.Activo, 
    &ru.FechaModificacion, 
    &ru.FechaCreacion,
)

    if err != nil {
        respondJSON(w, 500, map[string]string{"error": "Error al insertar la rutina: " + err.Error()})
        return
    }

    respondJSON(w, 201, ru)
}
// UpdateRutina actualiza una rutina por id_usuario y nombre
func UpdateRutina(w http.ResponseWriter, r *http.Request) {
	idUsuario := mux.Vars(r)["id_usuario"]
	var ru models.Rutina
	if err := json.NewDecoder(r.Body).Decode(&ru); err != nil {
		respondJSON(w, 400, map[string]string{"error": "JSON inválido"})
		return
	}

	_, err := config.DB.Exec(
		`UPDATE rutinas.rutinas 
          SET nombre=$1, descripcion=$2, objetivo=$3, dia=$4, activo=$5
          WHERE id_usuario=$6`,
		ru.Nombre, ru.Descripcion, ru.Objetivo, ru.Dia, ru.Activo, idUsuario,
	)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Rutina actualizada correctamente"})
}
// DeleteRutina elimina las rutinas 
func DeleteRutina(w http.ResponseWriter, r *http.Request) {
	idUsuario := mux.Vars(r)["id_usuario"]

	_, err := config.DB.Exec("DELETE FROM rutinas.rutinas WHERE id_usuario=$1", idUsuario)
	if err != nil {
		respondJSON(w, 500, map[string]string{"error": err.Error()})
		return
	}

	respondJSON(w, 200, map[string]string{"message": "Rutina(s) eliminada(s) correctamente"})
}