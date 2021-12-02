package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"time"
)

const CACHED_VALUE_TTL = 60
const REDIS_WORKPLACE_ID_TEMPLATE = "WORKPLACE:%d"

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(host string, port int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port), // "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	if pong != "PONG" {
		return nil, errors.New("Redis server is not responding")
	}

	return &RedisClient{
		Client: client,
	}, nil
}

func (c *RedisClient) CacheWorkplace(workplace model.Workplace) error {
	eventJson, err := json.Marshal(workplace)
	if err != nil {
		return err
	}

	err = c.Client.Set(fmt.Sprintf(REDIS_WORKPLACE_ID_TEMPLATE, workplace.ID), eventJson, CACHED_VALUE_TTL*time.Second).Err()

	return err
}

func (c *RedisClient) GetWorkplaceFromCache(workplaceId uint64) (*model.Workplace, bool, error) {
	workplaceJson, err := c.Client.Get(fmt.Sprintf(REDIS_WORKPLACE_ID_TEMPLATE, workplaceId)).Result()
	if err == redis.Nil {
		return nil, false, nil
	}

	if err != nil {
		return nil, false, err
	}

	workplace := new(model.Workplace)
	err = json.Unmarshal([]byte(workplaceJson), &workplace)
	if err != nil {
		return nil, false, err
	}

	return workplace, true, nil
}

func (c *RedisClient) DeleteCachedWorkplace(workplaceId uint64) error {
	err := c.Client.Del(fmt.Sprintf(REDIS_WORKPLACE_ID_TEMPLATE, workplaceId)).Err()

	return err
}
