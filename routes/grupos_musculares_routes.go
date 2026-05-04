package routes

import (
	"api_rutinas_gym/controllers"
	"github.com/gorilla/mux"
)
func RegisterGruposMuscularesRoutes(r *mux.Router) {
	r.HandleFunc("/grupos_musculares", controllers.GetAllGruposMusculares).Methods("GET")
	r.HandleFunc("/grupos_musculares/{id}", controllers.GetGrupoMuscularByID).Methods("GET")
	r.HandleFunc("/grupos_musculares", controllers.CreateGrupoMuscular).Methods("POST")
	r.HandleFunc("/grupos_musculares/{id}", controllers.UpdateGrupoMuscular).Methods("PUT")
	r.HandleFunc("/grupos_musculares/{id}", controllers.DeleteGrupoMuscular).Methods("DELETE")
}