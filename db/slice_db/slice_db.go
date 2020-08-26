package slice_db

type Db struct {
	numbers []int
}

func New() *Db {
	return &Db{}
}

// Create new item
func (db *Db) Add(num int) string {
	for _, item := range db.numbers {
		if item == num {
			return "Number already exist in the database"
		}
	}
	db.numbers = append(db.numbers, num)
	return "Number successfully added to the database"
}

// Get all items
func (db *Db) GetAll() []int {
	return db.numbers
}

// Get specific item
func (db *Db) GetOne(num int) interface{} {
	// If not found return not found
	for _, item := range db.numbers {
		if item == num {
			return item
		}
	}
	return "Number is not found"
}

// Delete specific item
func (db *Db) DeleteOne(num int) interface{} {
	newArr := make([]int, len(db.numbers))
	copy(newArr, db.numbers)

	for i := 0; i < len(newArr); i++ {
		if newArr[i] == num {
			newArr = append(newArr[:i], newArr[i+1:]...)
			i--
			return newArr
		}
	}
	return "Number is not found"
}
