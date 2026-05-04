# API Rutinas Gym - Go CRUD

API REST en Go para el esquema `rutinas` del proyecto de gimnasios.

## Tablas cubiertas
- `rutinas.grupos_musculares`
- `rutinas.ejercicios`
- `rutinas.rutinas`

## Configuración

1. Abre `config/db.go` y actualiza las credenciales:
```go
host     := "localhost"
port     := 5432
user     := "postgres"
password := "postgres"  
dbname   := "thehousefit_db"  
schema   := "rutinas"            
```

2. Instala dependencias y ejecuta:
```bash
go mod tidy
go run main.go
```

El servidor corre en: `http://localhost:8082`

---
**Body POST/PUT:**
```json
{
  "nombre": "Pecho",
  "descripcion": "Músculos del pectoral",
  "activo": true
}
```

---
**Body POST/PUT:**
```json
{
  "grupo_muscular_id": 1,
  "nombre": "Press de banca",
  "descripcion_corta": "Ejercicio básico de pecho",
  "descripcion_larga": "Ejercicio compuesto que trabaja principalmente el pectoral mayor",
  "posicion_inicial": "Tumbado en el banco con la barra a la altura del pecho",
  "ejecucion": "Empuja la barra hacia arriba extendiendo los brazos",
  "consejos": "Mantén los pies apoyados en el suelo",
  "nivel": "intermedio",
  "activo": true
}
```
**Body POST/PUT:**
```json
{
  "id_usuario": 1,
  "nombre": "Rutina Fuerza ",
  "descripcion": "Rutina de fuerza para principiantes",
  "objetivo": "ganar fuerza",
  "activo": true
}
```