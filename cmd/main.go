package main

import (
	"log"
	"net/http"

	"github.com/Jimb055/parking-status/internal/handler"
)

func main() {
	http.HandleFunc("/api/parking/disponibles", handler.GetAvailableParkings)

	log.Println("🚗 ParkingStatus service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
