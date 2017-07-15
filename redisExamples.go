package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

var client *redis.Client

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	return client
}

func ExampleClient(client *redis.Client) {
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	f := []float64{1.0, 2.0}
	data, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}

	errr := client.Set("key", data, 0).Err()
	if errr != nil {
		panic(errr)
	}
	val, _ := client.Get("key").Result()

	fmt.Println("key", val) // Making sure val is over-written

	val2, err := client.Get("key2").Result() // Key should not exist
	if err == redis.Nil {
		fmt.Println("Key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	client.SAdd("keyzo", "foo bar", "fizz buzz", "r00ty") // function overloading to add multiple strings
}

func main() {
	fmt.Println("go")
	client = NewClient()
	ExampleClient(client)
}
