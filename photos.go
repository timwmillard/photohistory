package photohistory

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Photo of the location
type Photo struct {
	ID          uuid.UUID `db:"id" json:"id"`
	TimeTaken   time.Time `db:"time_taken" json:"time_taken"`
	LocationID  uuid.UUID `db:"location_id" json:"location_id"`
	URL         string    `db:"url" json:"url"`
	SubmittedBy string    `db:"submitted_by" json:"submitted_by"`
	Latitude    float64   `db:"latitude" json:"latitude"`
	Longitude   float64   `db:"longitude" json:"longitude"`
	Elevation   float64   `db:"elevation" json:"elevation"`
}

// PhotosStore is to store photos
type PhotosStore interface {
	List() ([]*Location, error)
	Create(l *Location) (*Location, error)
	Update(id uuid.UUID, l *Location) (*Location, error)
	Get(id uuid.UUID) (*Location, error)
	Delete(id uuid.UUID) error
}

// Common Errors
var (
	ErrPhotoNotFound = errors.New("photo not found")
)
