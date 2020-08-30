package slice_db

import "errors"

type SliceDb struct {
	numbers []int
}

func New() *SliceDb {
	return &SliceDb{
		numbers: []int{},
	}
}

// Create new item
func (db *SliceDb) Add(num int) bool {
	for _, item := range db.numbers {
		if item == num {
			return false
		}
	}
	db.numbers = append(db.numbers, num)
	return true
}

// Get all items
func (db *SliceDb) GetAll() []int {
	return db.numbers
}

// Get specific item
func (db *SliceDb) GetOne(num int) (int, error) {
	for _, item := range db.numbers {
		if item == num {
			return item, nil
		}
	}
	return 0, errors.New("Number is not found")
}

// Update specific item
// func (db *SliceDb) Update(num int, newNum int) error {
// 	for _, item := range db.numbers {
// 		if item == newNum {
// 			// New number already exist
// 			return errors.New("New number already exists in a database")
// 		}
// 	}

// 	for i := 0; i < len(db.numbers); i++ {
// 		if db.numbers[i] == num {
// 			db.numbers = append(db.numbers[:i], db.numbers[i+1:]...)
// 			db.numbers = append(db.numbers, newNum)
// 			return nil
// 		}
// 	}
// 	return errors.New("Searched number is not found in a database")
// }

// Delete specific item
func (db *SliceDb) Delete(num int) bool {
	for i := 0; i < len(db.numbers); i++ {
		if db.numbers[i] == num {
			db.numbers = append(db.numbers[:i], db.numbers[i+1:]...)
			return true
		}
	}
	return false
}
