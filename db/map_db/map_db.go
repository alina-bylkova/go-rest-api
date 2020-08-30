package map_db

import "errors"

type MapDb struct {
	numbers map[int]bool
}

func New() *MapDb {
	return &MapDb{
		numbers: map[int]bool{},
	}
}

// Create new item
func (db *MapDb) Add(num int) bool {
	for k := range db.numbers {
		if k == num {
			return false
		}
	}
	db.numbers[num] = true
	return true
}

// Get all items
func (db *MapDb) GetAll() []int {
	arr := []int{}
	for k := range db.numbers {
		arr = append(arr, k)
	}
	return arr
}

// Get specific item
func (db *MapDb) GetOne(num int) (int, error) {
	// If not found return not found
	for k := range db.numbers {
		if k == num {
			return k, nil
		}
	}
	return 0, errors.New("Number is not found")
}

// Update specific item
// func (db *MapDb) Update(num int, newNum int) error {
// 	for k := range db.numbers {
// 		if k == newNum {
// 			// New number already exist
// 			return errors.New("New number already exists in a database")
// 		}
// 	}

// 	for k := range db.numbers {
// 		if k == num {
// 			delete(db.numbers, num)
// 			db.numbers[newNum] = true
// 			return nil
// 		}
// 	}
// 	return errors.New("Searched number is not found in a database")
// }

// Delete specific item
func (db *MapDb) Delete(num int) bool {
	for k := range db.numbers {
		if k == num {
			delete(db.numbers, num)
			return true
		}
	}
	return false
}
