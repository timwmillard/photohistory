package postgres

import (
	"photohistory"

	"github.com/google/uuid"

	// Need MySQL driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Need PostgreSQL driver
)

// PhotosStore -
type PhotosStore struct {
	db *sqlx.DB
}

// List all photos
func (ps *PhotosStore) List() ([]*photohistory.Photo, error) {
	var p photohistory.Photo
	if err := ps.db.Get(&p, `SELECT * FROM photos`); err != nil {
		return nil, err
	}
	return p, nil
}

// Create a single photo
func (ps *PhotosStore) Create(p *photohistory.Photo) (*photohistory.Photo, error) {
	q := `INSERT INTO photos
			(id, time_taken, location_id, url, submitted_by, latitude, longitude, elevation)
			VALUE ((UUID_TO_BIN(UUID()), :time_taken, :location_id, :url, :submitted_by, :latitude, :longitude, :elevation)`
	result, err := ps.db.NamedExec(q, p)
	if err != nil {
		return nil, err
	}

	// Update the competitor wit the new ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var newPhoto photohistory.Photo
	q = `SELECT id, time_taken, location_id, url, submitted_by, latitude, longitude, elevation
			FROM photos
			WHERE id=?`
	err = ps.db.Get(&newPhoto, q, id)
	if err != nil {
		return nil, err
	}

	return &newPhoto, nil
}

// Update a photo
func (ps *PhotosStore) Update(p *photohistory.Photo) error {
	q := `UPDATE photos
			(id, time_taken, location_id, url, submitted_by, latitude, longitude, elevation)
			VALUE :id, :time_taken, :location_id, :url, :submitted_by, :latitude, :longitude, :elevation)`
	_, err := ps.db.NamedExec(q, p)
	if err != nil {
		return err
	}
	return nil
}

// Get a single photo
func (ps *PhotosStore) Get(id uuid.UUID) (*photohistory.Photo, error) {
	var p photohistory.Photo
	if err := ps.db.Get(&p, `SELECT * FROM photos WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return p, nil
}

// Delete a photo
func (ps *PhotosStore) Delete(id uuid.UUID) error {
	if _, err := ps.db.Exec(`DELETE FROM photos WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
