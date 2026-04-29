package routes

import (
	"api_rutinas_gym/controllers"
	"github.com/gorilla/mux"
)
func RegisterRutinasRoutes(r *mux.Router) {
	r.HandleFunc("/rutinas", controllers.GetAllRutinas).Methods("GET")
	r.HandleFunc("/rutinas/{id_usuario}", controllers.GetRutinaByUsuario).Methods("GET")
	r.HandleFunc("/rutinas", controllers.CreateRutina).Methods("POST")
	r.HandleFunc("/rutinas/{id_usuario}", controllers.UpdateRutina).Methods("PUT")
	r.HandleFunc("/rutinas/{id_usuario}", controllers.DeleteRutina).Methods("DELETE")
}