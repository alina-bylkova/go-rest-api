package slice_db

import (
	"testing"
)

func TestAdd(t *testing.T) {
	db := New()
	db.Add(5)
	if len(db.numbers) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	db := New()
	db.Add(5)
	results := db.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetOne(t *testing.T) {
	db := New()
	db.Add(5)
	db.Add(10)
	result := db.GetOne(5)
	if result == "Number is not found" {
		t.Errorf("Number is not found")
	}
}

func TestDeleteOne(t *testing.T) {
	db := New()
	db.Add(5)
	db.Add(10)
	results, _ := db.DeleteOne(5).([]int)
	if len(results) != 1 {
		t.Errorf("Item was not deleted")
	}
}
