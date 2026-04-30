package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// ConnectDB establece conexión con PostgreSQL
func ConnectDB() {
	host     := "localhost"
	port     := 5432
	user     := "postgres"
	password := "postgres" // 
	dbname   := "thehousefit_db" // 
	schema   := "rutinas"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error al abrir conexión: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("No se puede conectar a la BD: ", err)
	}

	fmt.Println("Conexión a base de datos exitosa")
	fmt.Println("Conectado a la DB:", dbname, "| Esquema:", schema)
	DB = db
}