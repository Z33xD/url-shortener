package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

/*
// Setting up wrappers around Redis, that will be used as an interface for persisting and retrieving the application's data mapping.
*/

// struct wrapper around the raw redis client
type StorageService struct {
	redisClient *redis.Client
}

var storeService = &StorageService{}
var ctx = context.Background()

// an LRU config would be too much for now, so I'm just gonna set an expiration timer
const CacheDuration = 6 * time.Hour

/*
// Initialising the store service (Redis client, in this project's case)
*/

func InitialiseStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis' Docker container's address
		Password: "",
		DB:       0, // Default database
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error initialising Redis. Error: %v", err))
	}

	fmt.Printf("\nRedis started successfully. Pong message: %s", pong)
	storeService.redisClient = redisClient
	return storeService
}

/*
// To save the mapping between the originalUrl, and the generated shortUrl
*/
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()

	if err != nil {
		panic(fmt.Sprintf(
			`Failed while saving the key URL.
			Error: %v
			Short URL: %v
			Original URL: %v`, err, shortUrl, originalUrl,
		))
	}
}

/*
// To retrieve the initial longURL once the shortUrl is provided.
// (Used when users will be calling the shortLink in the URL)
*/
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()

	if err != nil {
		panic(fmt.Sprintf(
			`Failed while retrieving the initial URL
			Error: %v
			Short URL: %v`, err, shortUrl,
		))
	}

	return result
}
