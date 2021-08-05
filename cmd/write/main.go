package main

import (
	"log"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{})

	deleteResult := client.Del("stream:foo")

	if deleteResult.Err() != nil {
		log.Fatalf("Could not clean up stream data: %v", deleteResult.Err())
	}

	for i := 0; i < 200_000; i++ {
		values := make(map[string]interface{})
		values["my-value"] = i

		args := redis.XAddArgs{
			Stream:       "stream:foo",
			MaxLen:       200_000,
			MaxLenApprox: 200_000,
			ID:           "*",
			Values:       values,
		}

		result := client.XAdd(&args)

		if result.Err() != nil {
			log.Fatalf("write failed: %v", result.Err())
		}
	}

	streamLength := client.XLen("stream:foo").Val()

	log.Println("stream has entries", streamLength)
}
