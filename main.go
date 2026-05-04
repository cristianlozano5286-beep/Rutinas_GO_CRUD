package main

import (
	"log"
	"net/http"

	"api_rutinas_gym/config"
	"api_rutinas_gym/routes"

	"github.com/gorilla/mux"
)
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	config.ConnectDB() // Conectar DB

	r := mux.NewRouter()

	// Registrar rutas del esquema rutinas
	routes.RegisterGruposMuscularesRoutes(r)
	routes.RegisterEjerciciosRoutes(r)
	routes.RegisterRutinasRoutes(r)

	log.Println("Servidor corriendo en el puerto 8082")
	http.ListenAndServe(":8082", enableCORS(r))
}