package postgres

import (
	"photohistory"

	"github.com/google/uuid"

	// Need MySQL driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Need PostgreSQL driver
)

// LocationsStore -
type LocationsStore struct {
	db *sqlx.DB
}

// List all locations
func (ls *LocationsStore) List() ([]*photohistory.Location, error) {
	return nil, nil
}

// Create a singel location
func (ls *LocationsStore) Create(l *photohistory.Location) (*photohistory.Location, error) {
	q := `INSERT INTO locations
			(id, alias, name, latitude, longitude, elevation)
			VALUE (UUID_TO_BIN(UUID()), :alias, :name, :latitude, :longitude, :elevation)`
	result, err := ls.db.NamedExec(q, l)
	if err != nil {
		return nil, err
	}

	// Update the competitor wit the new ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var newLocation photohistory.Location
	q = `SELECT id, alias, name, latitude, longitude, elevation
			FROM locations
			WHERE id=?`
	err = ls.db.Get(&newLocation, q, id)
	if err != nil {
		return nil, err
	}

	return &newLocation, nil
}

// Update a location
func (ls *LocationsStore) Update(l *photohistory.Location) error {
	q := `UPDATE locations
			SET (id, alias, name, latitude, longitude, elevation)
			VALUE :id, :alias, :name, :latitude, :longitude, :elevation)`
	_, err := ls.db.NamedExec(q, l)
	if err != nil {
		return err
	}
	return nil
}

// Get a single location
func (ls *LocationsStore) Get(id uuid.UUID) (*photohistory.Location, error) {
	var l photohistory.Location
	if err := ls.db.Get(&l, `SELECT * FROM locations WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return l, nil
}

// Delete a location
func (ls *LocationsStore) Delete(id uuid.UUID) error {
	if _, err := ls.db.Exec(`DELETE FROM locations WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}

// NearbyLocation(p *Location, radiusKm float32) ([]Location, error)
