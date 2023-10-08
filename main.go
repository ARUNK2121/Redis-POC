package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	//ping to check if the connection is working properly
	result, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(result)

	// Set a key-value pair
	if err := client.Set("mykey", "this is the new value for mykey", 0).Err(); err != nil {
		panic(err)
	}

	// Retrieve a value for a given key
	val, err := client.Get("mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey", val)
}
