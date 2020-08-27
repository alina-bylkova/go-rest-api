package slice_db

type Db struct {
	Numbers []int
}

func New() *Db {
	// return &Db{
	// 	Numbers: []int{5, 10},
	// }
	return &Db{}
}

// Create new item
func (db *Db) Add(num int) interface{} {
	for _, item := range db.Numbers {
		if item == num {
			return "Number already exist in the database"
		}
	}
	db.Numbers = append(db.Numbers, num)
	return db.Numbers
}

// Get all items
func (db *Db) GetAll() []int {
	return db.Numbers
}

// Get specific item
func (db *Db) GetOne(num int) interface{} {
	// If not found return not found
	for _, item := range db.Numbers {
		if item == num {
			return item
		}
	}
	return "Number is not found"
}

// Delete specific item
func (db *Db) Delete(num int) interface{} {
	for i := 0; i < len(db.Numbers); i++ {
		if db.Numbers[i] == num {
			db.Numbers = append(db.Numbers[:i], db.Numbers[i+1:]...)
			i--
			return db.Numbers
		}
	}
	return "Number is not found"
}
