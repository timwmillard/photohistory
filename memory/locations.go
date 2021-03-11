package memory

import (
	"photohistory"

	"github.com/google/uuid"
)

// LocationsStore in memory store of locations
type LocationsStore struct {
	data map[uuid.UUID]*photohistory.Location
}

// NewLocationsStore create a new location store in memeory and creates the internal map.
func NewLocationsStore() *LocationsStore {
	data := make(map[uuid.UUID]*photohistory.Location)
	store := &LocationsStore{data: data}
	return store
}

// List all locations from the store.
func (ls *LocationsStore) List() ([]*photohistory.Location, error) {
	var list []*photohistory.Location
	for _, location := range ls.data {
		list = append(list, location)
	}
	return list, nil
}

// Get a single location from the store by id.
func (ls *LocationsStore) Get(id uuid.UUID) (*photohistory.Location, error) {
	location, ok := ls.data[id]
	if !ok {
		return nil, photohistory.ErrLocationNotFound
	}
	return location, nil
}

// Create -
func (ls *LocationsStore) Create(l *photohistory.Location) (*photohistory.Location, error) {
	id := uuid.New()
	l.ID = id
	ls.data[id] = l
	return l, nil
}

// Update -
func (ls *LocationsStore) Update(id uuid.UUID, l *photohistory.Location) (*photohistory.Location, error) {
	_, ok := ls.data[id]
	if !ok {
		return nil, photohistory.ErrLocationNotFound
	}
	l.ID = id
	ls.data[id] = l
	return l, nil
}

// Delete -
func (ls *LocationsStore) Delete(id uuid.UUID) error {
	delete(ls.data, id)
	return nil
}
