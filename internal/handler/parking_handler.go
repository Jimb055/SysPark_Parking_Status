package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Jimb055/parking-status/internal/service"
)

func GetAvailableParkings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	parkings := service.FetchAvailableParkings()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parkings)
}
