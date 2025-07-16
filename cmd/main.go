package main

import (
	"log"
	"net/http"

	"github.com/Jimb055/parking-status/internal/database"
	"github.com/Jimb055/parking-status/internal/handler"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar archivo .env si existe (útil en desarrollo local)
	_ = godotenv.Load()

	// Inicializar conexión con la base de datos
	database.ConnectPostgres()

	// Registrar endpoints
	http.HandleFunc("/api/parking/disponibles", handler.GetAvailableParkings)

	log.Println("🚗 ParkingStatus service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
