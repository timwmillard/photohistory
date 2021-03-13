package postgres

import (
	"photohistory"

	"github.com/google/uuid"

	// Need MySQL driver
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Need PostgreSQL driver
)

// UsersStore -
type UsersStore struct {
	db *sqlx.DB
}

// List all photos
func (us *UsersStore) List() ([]*photohistory.User, error) {
	var u photohistory.User
	if err := us.db.Get(&u, `SELECT * FROM users`); err != nil {
		return nil, err
	}
	return u, nil
}

// Create a single photo
func (us *UsersStore) Create(u *photohistory.User) (*photohistory.User, error) {
	q := `INSERT INTO users
			(id, time_taken, location_id, url, submitted_by, latitude, longitude, elevation)
			VALUE ((UUID_TO_BIN(UUID()), :time_taken, :location_id, :url, :submitted_by, :latitude, :longitude, :elevation)`
	result, err := us.db.NamedExec(q, u)
	if err != nil {
		return nil, err
	}

	// Update the competitor with the new ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	var newUser photohistory.User
	q = `SELECT id, time_taken, location_id, url, submitted_by, latitude, longitude, elevation
			FROM users
			WHERE id=?`
	err = us.db.Get(&newUser, q, id)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

// Update a photo
func (us *UsersStore) Update(u *photohistory.User) error {
	q := `UPDATE users
	        (id, username, first_name, last_name, email, phone)
	        VALUE :id, :username, :first_name, :last_name, :email, :phone`
	_, err := us.db.NamedExec(q, u)
	if err != nil {
		return err
	}
	return nil
}

// Get a single photo
func (us *UsersStore) Get(id uuid.UUID) (*photohistory.User, error) {
	var u photohistory.User
	if err := us.db.Get(&u, `SELECT * FROM users WHERE id = $1`, id); err != nil {
		return nil, err
	}
	return u, nil
}

// Delete a photo
func (us *UsersStore) Delete(id uuid.UUID) error {
	if _, err := us.db.Exec(`DELETE FROM users WHERE id = $1`, id); err != nil {
		return err
	}
	return nil
}
