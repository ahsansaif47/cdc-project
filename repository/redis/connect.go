package redis

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/ahsansaif47/cdc-app/config"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Store(ctx context.Context, key string, value any, expirationTime time.Duration) error
	Retrieve(ctx context.Context, key string) (any, error)
}

type cache struct {
	Client *redis.Client
}

func NewCache() Cache {
	return &cache{
		Client: connect(),
	}
}

func connect() *redis.Client {
	c := config.GetConfig()
	opt, err := redis.ParseURL(c.RedisUrl)
	if err != nil {
		log.Fatalf("error parsing redis config: %v", err)
	}

	client := redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("error connecting to redis: %v", err)
	}
	return client
}

func (c *cache) Store(ctx context.Context, key string, value any, expirationTime time.Duration) error {
	storeValue, err := json.Marshal(value)
	if err != nil {
		log.Println("error marshalling value: ", err)
		return err
	}

	return c.Client.Set(ctx, key, storeValue, expirationTime).Err()
}

func (c *cache) Retrieve(ctx context.Context, key string) (any, error) {
	value, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		log.Println("error retrieving value: ", err)
		return nil, err
	}

	return value, nil
}
