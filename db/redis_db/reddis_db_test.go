package redis_db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	db := New()
	result := db.Add(20)
	if !result {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	db := New()
	db.Add(20)
	results := db.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetOne(t *testing.T) {
	db := New()
	db.Add(5)
	db.Add(10)
	result, err := db.GetOne(5)

	assert.Nil(t, err)
	assert.Equal(t, 5, result, "they should be equal")

	_, err = db.GetOne(6)
	assert.NotNil(t, err, "Should be an error as the object is not found")
}

func TestDelete(t *testing.T) {
	db := New()
	db.Add(5)
	db.Add(10)
	// result := db.Delete(5)
	// if !result {
	// 	t.Errorf("Item was not deleted")
	// }
	db.Delete(5)
	resultGet, _ := db.GetOne(5)
	if resultGet == 5 {
		t.Errorf("Item was not deleted")
	}
}
