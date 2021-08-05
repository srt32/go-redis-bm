package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{})

	start := time.Now()

	result := client.XRange("stream:foo", "-", "+")
	count := 0
	for _, element := range result.Val() {
		_, err := strconv.Atoi(element.Values["my-value"].(string))

		if err != nil {
			log.Fatal("counting failed", err)
		}

		count += 1
	}

	timeDiff := time.Now().Sub(start)

	fmt.Printf("total elements read: %v that took %v", count, timeDiff)
}
