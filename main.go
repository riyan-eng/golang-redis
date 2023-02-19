package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func Rds() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func main() {

	ctx := context.Background()
	// if err := Rds().Set(ctx, "nama", "riyan", 0); err != nil {
	// 	fmt.Println(err)
	// }

	val, err := Rds().Get(ctx, "nama").Result()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
