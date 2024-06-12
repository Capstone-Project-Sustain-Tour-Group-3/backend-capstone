package redis

import (
	"context"
	"errors"
	"time"

	"capstone/errorHandlers"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	UpdateChatHistory(key string, value string) error
	GetChatHistory(key string) (*[]string, error)
	DeleteChatHistory(key string) error
	SetRecommendedDestinationsIds(key string, values []string) error
	GetRecommendedDestinationsIds(key string) (*[]string, error)
}

type redisClient struct {
	client *redis.Client
}

func NewRedisClient(addr, username, password string) *redisClient {
	return &redisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     addr,
			Username: username,
			Password: password,
		}),
	}
}

func (r *redisClient) UpdateChatHistory(key string, value string) error {
	ctx := context.Background()

	err := r.client.LPush(ctx, key, value).Err()
	if err != nil {
		return err
	}

	err = r.client.Expire(ctx, key, 6*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisClient) GetChatHistory(key string) (*[]string, error) {
	ctx := context.Background()

	var maxHistory int64 = 10
	history, err := r.client.LRange(ctx, key, 0, maxHistory-1).Result()
	if err != nil {
		return nil, err
	}

	if len(history) == 0 {
		return nil, errors.New("chat history not found")
	}

	return &history, nil
}

func (r *redisClient) DeleteChatHistory(key string) error {
	ctx := context.Background()

	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return &errorHandlers.InternalServerError{
			Message: "failed to delete chat history",
		}
	}

	return nil
}

func (r *redisClient) GetRecommendedDestinationsIds(key string) (*[]string, error) {
	ctx := context.Background()

	history, err := r.client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	return &history, nil
}

func (r *redisClient) SetRecommendedDestinationsIds(key string, values []string) error {
	ctx := context.Background()

	err := r.client.RPush(ctx, key, values).Err()
	if err != nil {
		return err
	}

	return nil
}
