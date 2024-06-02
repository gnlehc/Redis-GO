package main

import (
	"context"
	"encoding/json"
	"fmt"
	"redis-go/model"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping the Redis server
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to ping Redis server: %s\n", err.Error())
		return
	}
	fmt.Printf("Ping response: %s\n", ping)

	// Set a value in Redis
	err = client.Set(context.Background(), "name", "lotso", 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the Redis instance: %s\n", err.Error())
		return
	}
	fmt.Println("Value 'lotso' set in Redis")

	// Get the value from Redis
	val, err := client.Get(context.Background(), "name").Result()
	if err != nil {
		fmt.Printf("Failed to get value from Redis: %s\n", err.Error())
		return
	}
	fmt.Printf("Value from Redis: %s\n", val)

	// model payload
	jsonString, err := json.Marshal(&model.Person{
		Name:  "Chelsea",
		Age:   19,
		Email: "Chelseang40@gmail.com",
	})
	if err != nil {
		fmt.Printf("Failed to marshal: %s", err.Error())
		return
	}

	// Set a value in Redis
	err = client.Set(context.Background(), "person", jsonString, 0).Err()
	if err != nil {
		fmt.Printf("Failed to set value in the Redis instance: %s\n", err.Error())
		return
	}
	fmt.Printf("Value %s set in Redis", jsonString)

	// Get the value from Redis
	vals, err := client.Get(context.Background(), "person").Result()
	if err != nil {
		fmt.Printf("Failed to get value from Redis: %s\n", err.Error())
		return
	}
	fmt.Println()
	fmt.Printf("Value from Redis: %s\n", vals)
}
