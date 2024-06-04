package reservation

import (
	"fmt"
	"time"
)

type Reservation struct {
	ID        int
	Name      string
	StartTime time.Time
	EndTime   time.Time
}

var reservations = make(map[int]Reservation)
var lastID int

func addReservation(name string, start, end time.Time) int {
	lastID++
	reservations[lastID] = Reservation{
		ID:        lastID,
		Name:      name,
		StartTime: start,
		EndTime:   end,
	}
	return lastID
}

func removeReservation(id int) error {
	if _, exists := reservations[id]; exists {
		delete(reservations, id)
		return nil
	}
	return fmt.Errorf("No reservation found with ID %d", id)
}

func findReservationByName(name string) []Reservation {
	var results []Reservation
	for _, res := range reservations {
		if res.Name == name {
			results = append(results, res)
		}
	}
	return results
}

func main() {

	start, _ := time.Parse("2006-01-02 15:04", "2023-03-25 14:00")
	end, _ := time.Parse("2006-01-02 15:04", "2023-03-25 16:00")
	id := addReservation("John Doe", start, end)
	fmt.Println("Added reservation:", reservations[id])

	searchResults := findReservationByName("John Doe")
	fmt.Println("Search results:", searchResults)

	err := removeReservation(id)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Reservation removed successfully.")
	}
}
