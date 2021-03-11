package memory

import (
	"photohistory"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCreateAndGet(t *testing.T) {
	store := NewLocationsStore()

	want, err := store.Create(&loc1)
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}

	got, err := store.Get(want.ID)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got a different location: got %v, want %v", got, want)
	}

}

func TestCreateAndList(t *testing.T) {
	store := NewLocationsStore()

	_, err := store.Create(&loc1)
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}
	_, err = store.Create(&loc2)
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}

	locations, err := store.List()
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}

	if len(locations) != 2 {
		t.Errorf("should be 2 locations but got %v locations", len(locations))
	}
}

func TestGetLocationNotFound(t *testing.T) {
	store := NewLocationsStore()
	id := uuid.New()
	_, err := store.Get(id)
	if err != photohistory.ErrLocationNotFound {
		t.Error("should have got ErrLocationNotFound")
	}
}

func TestCreateAndUpdate(t *testing.T) {
	store := NewLocationsStore()

	l, err := store.Create(&loc1)
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}

	want := &loc2
	got, err := store.Update(l.ID, want)
	if err != nil {
		t.Fatalf("unable to update location: %v", err)
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got a different location: got %v, want %v", got, want)
	}

}

func TestUpdateLocationNotFound(t *testing.T) {
	store := NewLocationsStore()
	id := uuid.New()
	_, err := store.Update(id, &loc1)
	if err != photohistory.ErrLocationNotFound {
		t.Error("should have got ErrLocationNotFound")
	}
}

func TestCreateAndDelete(t *testing.T) {
	store := NewLocationsStore()

	l, err := store.Create(&loc1)
	if err != nil {
		t.Fatalf("unable to create location: %v", err)
	}

	err = store.Delete(l.ID)
	if err != nil {
		t.Fatalf("unable to delete location: %v", err)
	}

	locations, err := store.List()
	if err != nil {
		t.Fatalf("unable to list locations: %v", err)
	}

	if len(locations) != 0 {
		t.Errorf("location list should be empty but has %d locations", len(locations))
	}
}

var (
	loc1 = photohistory.Location{
		Alias:     "bendigo",
		Name:      "Bendigo Park",
		Latitude:  36.7570,
		Longitude: 144.2794,
		Elevation: 213,
	}
	loc2 = photohistory.Location{
		Alias:     "koondrook",
		Name:      "Koondrook Park",
		Latitude:  35.6333,
		Longitude: 144.1167,
		Elevation: 78,
	}
)
