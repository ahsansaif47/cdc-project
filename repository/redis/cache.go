package redis

import (
	"context"
	"time"
)

type ICacheRepository interface {
	StoreOTP(email, otp string, expTime time.Duration) error
	RetrieveOTP(email string) (string, error)
}

type CacheRepository struct {
	cache Cache
}

func NewUserCache(cache Cache) ICacheRepository {
	return &CacheRepository{cache: cache}
}

func (r *CacheRepository) StoreOTP(email, otp string, expTime time.Duration) error {
	return r.cache.Store(context.Background(), email, otp, expTime)
}

func (r *CacheRepository) RetrieveOTP(email string) (string, error) {
	value, err := r.cache.Retrieve(context.Background(), email)
	if err != nil {
		return "", err
	}

	otp, ok := value.(string)
	if !ok {
		return "", nil
	}

	return otp, nil
}
