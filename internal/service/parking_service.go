package service

import (
	"log"

	"github.com/Jimb055/parking-status/internal/database"
)

type ParkingSlot struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
	Status   string `json:"status"` // "available" o "occupied"
}

func FetchAvailableParkings() []ParkingSlot {
	db := database.GetDB()

	rows, err := db.Query("SELECT id, ubicacion, disponible FROM espacios")
	if err != nil {
		log.Println("❌ Error al consultar la base de datos:", err)
		return nil
	}
	defer rows.Close()

	var results []ParkingSlot
	for rows.Next() {
		var p ParkingSlot
		var disponible bool

		err := rows.Scan(&p.ID, &p.Location, &disponible)
		if err != nil {
			log.Println("❌ Error al escanear fila:", err)
			continue
		}

		if disponible {
			p.Status = "available"
		} else {
			p.Status = "occupied"
		}

		results = append(results, p)
	}

	return results
}
