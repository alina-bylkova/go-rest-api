package db

// Interfaces can be used as function parameters in helper functions like NumbersGet, NumbersPost, etc.
type Getter interface {
	GetAll() []int
}

type Adder interface {
	Add(num int) bool
}
