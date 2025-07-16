package service

import "testing"

func TestFetchAvailableParkings(t *testing.T) {
	parkings := FetchAvailableParkings()

	if parkings == nil {
		t.Fatal("Expected a non-nil slice")
	}

	if len(parkings) == 0 {
		t.Fatal("Expected at least one parking slot")
	}

	for _, p := range parkings {
		if p.Status != "available" {
			t.Errorf("Expected status to be 'available', got '%s'", p.Status)
		}
	}
}
