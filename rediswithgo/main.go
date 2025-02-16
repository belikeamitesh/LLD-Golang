package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	fmt.Println("Redis with Go!")

	// NewClient returns a client to the Redis Server specified by Options.
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		//database to be selected after connecting server
		DB: 0,
	})

	ping,err := client.Ping(context.Background()).Result()
	if err!=nil{
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Successfully connected to redis Instance")
	fmt.Println("Pinged Succesfully with response: ",ping)

	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Occupation string  `json:"occupation"`
	}

	jsonString, err := json.Marshal(Person{
			Name: "Amitesh",
			Age: 25,
			Occupation: "Software Engineer",
		})
	if err!=nil{
		fmt.Println("Failed to marshal json : ",err.Error())
		return
	}

	// err = client.Set(context.Background(),"name","Amitesh",0).Err()
	err = client.Set(context.Background(),"person",jsonString,0).Err()
	if err!=nil{
		fmt.Println("Failed to set value in redis instance : ",err.Error())
		return
	}
	fmt.Println("Successfully set key value to redis Instance")


	val,err := client.Get(context.Background(),"person").Result()
	if err!=nil{
		fmt.Println("Failed to get value from redis instance : ",err.Error())
		return
	}
	fmt.Println("Value retrived from redis: ",val)
}