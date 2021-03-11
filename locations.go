package photohistory

import (
	"errors"

	"github.com/google/uuid"
)

// Location were photos can be taken
type Location struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Alias     string    `db:"alias" json:"alias"`
	Name      string    `db:"name" json:"name"`
	Latitude  float64   `db:"latitude" json:"latitude"`
	Longitude float64   `db:"longitude" json:"longitude"`
	Elevation float64   `db:"elevation" json:"elevation"`
}

// LocationStore is to store locations
type LocationStore interface {
	List() ([]*Location, error)
	Create(l *Location) (*Location, error)
	Update(id uuid.UUID, l *Location) (*Location, error)
	Get(id uuid.UUID) (*Location, error)
	Delete(id uuid.UUID) error
	// NearbyLocation(p *Location, radiusKm float32) ([]Location, error)
}

// Common Errors
var (
	ErrLocationNotFound = errors.New("location not found")
)
