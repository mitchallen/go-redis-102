package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	json, err := json.Marshal(User{Name: "Otto", Age: 45})
	if err != nil {
		fmt.Println(err)
	}

	TEST_KEY := "id123"

	err = client.Set(TEST_KEY, json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

	val, err := client.Get(TEST_KEY).Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
