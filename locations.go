package photohistory

import (
	"github.com/google/uuid"
)

// Location -
type Location struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Alias     string    `db:"alias" json:"alias"`
	Name      string    `db:"name" json:"name"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
	Elevation float64   `db:"elevation" json:"elevation"`
}

type LocationStore interface {
	List() ([]*Location, error)
	Create(l *Location) (*Location, error)
	Update(l *Location) error
	Get(id uuid.UUID) (*Location, error)
	Delete(id uuid.UUID) error
	// NearbyLocation(p *Location, radiusKm float32) ([]Location, error)
}
