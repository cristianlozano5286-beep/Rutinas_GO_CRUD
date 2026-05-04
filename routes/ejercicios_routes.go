package routes

import (
	"api_rutinas_gym/controllers"
	"github.com/gorilla/mux"
)
func RegisterEjerciciosRoutes(r *mux.Router) {
	r.HandleFunc("/ejercicios", controllers.GetAllEjercicios).Methods("GET")
	r.HandleFunc("/ejercicios/{id}", controllers.GetEjercicioByID).Methods("GET")
	r.HandleFunc("/ejercicios", controllers.CreateEjercicio).Methods("POST")
	r.HandleFunc("/ejercicios/{id}", controllers.UpdateEjercicio).Methods("PUT")
	r.HandleFunc("/ejercicios/{id}", controllers.DeleteEjercicio).Methods("DELETE")
}
