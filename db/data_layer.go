package db

// Interfaces can be used as function parameters in helper functions like NumbersGet, NumbersPost, etc.
type DataLayer interface {
	Add(int) bool
	GetAll() []int
	GetOne(int) (int, error)
	Delete(int) bool
	// Update(int, int) error
}
