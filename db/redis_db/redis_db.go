package redis_db

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisDb struct {
	rdb *redis.Client
	ctx context.Context
}

func New(url string) *RedisDb {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RedisDb{rdb, ctx}
}

// Create new item
func (db *RedisDb) Add(num int) bool {

	numString := strconv.Itoa(num)

	_, err := db.rdb.Get(db.ctx, numString).Result()
	if err != nil && err != redis.Nil {
		fmt.Printf("Internal error with redis: %v\n", err)
		return false
	}
	if err == nil {
		return false
	}

	if err := db.rdb.Set(db.ctx, numString, true, 0).Err(); err != nil {
		fmt.Printf("Internal error with redis: %v\n", err)
		return false
	}
	return true
}

// Get all items
func (db *RedisDb) GetAll() []int {
	val := db.rdb.Keys(db.ctx, "*")
	strs := val.Val()
	arr := make([]int, 0, len(strs))
	for _, str := range strs {
		num, err := strconv.Atoi(str)
		if err == nil {
			arr = append(arr, num)
		}
	}
	return arr
}

// Get specific item
func (db *RedisDb) GetOne(num int) (int, error) {
	numString := strconv.Itoa(num)

	_, err := db.rdb.Get(db.ctx, numString).Result()
	if err != nil && err != redis.Nil {
		fmt.Printf("Error: %v\n", err)
		return 0, errors.New("Internal error with redis")
	}
	if err != nil {
		return 0, errors.New("Number is not found")
	}
	return num, nil
}

// Delete specific item
func (db *RedisDb) Delete(num int) bool {
	numString := strconv.Itoa(num)

	val := db.rdb.Del(db.ctx, numString).Val()

	if val != 1 {
		return false
	}
	return true
}
