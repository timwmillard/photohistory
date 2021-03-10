package photohistory

import (
	"time"

	"github.com/google/uuid"
)

type Photo struct {
	ID          uuid.UUID `db:"id",json:"id"`
	TimeTaken   time.Time `db:"time_taken",json:"time_taken"`
	LocationID  uuid.UUID `db:"location_id",json:"location_id"`
	URL         string    `db:"url",json:"url"`
	SubmittedBy string    `db:"submitted_by",json:"submitted_by"`
	Latitude    float64   `db:"latitude",json:"latitude"`
	Longitude   float64   `db:"longitude",json:"longitude"`
	Elevation   float64   `db:"elevation",json:"elevation"`
}
