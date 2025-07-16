package service

type ParkingSlot struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
	Status   string `json:"status"` // e.g., "available", "occupied"
}

func FetchAvailableParkings() []ParkingSlot {
	// Mock data for demo
	return []ParkingSlot{
		{ID: 1, Location: "Zone A - 01", Status: "available"},
		{ID: 2, Location: "Zone A - 03", Status: "available"},
	}
}
