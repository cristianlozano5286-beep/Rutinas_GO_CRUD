package routes

import (
	"api_rutinas_gym/controllers"
	"github.com/gorilla/mux"
)
func RegisterGruposMuscularesRoutes(r *mux.Router) {
	r.HandleFunc("/grupos-musculares", controllers.GetAllGruposMusculares).Methods("GET")
	r.HandleFunc("/grupos-musculares/{id}", controllers.GetGrupoMuscularByID).Methods("GET")
	r.HandleFunc("/grupos-musculares", controllers.CreateGrupoMuscular).Methods("POST")
	r.HandleFunc("/grupos-musculares/{id}", controllers.UpdateGrupoMuscular).Methods("PUT")
	r.HandleFunc("/grupos-musculares/{id}", controllers.DeleteGrupoMuscular).Methods("DELETE")
}