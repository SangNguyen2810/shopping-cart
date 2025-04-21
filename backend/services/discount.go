package services

import (
	"context"
	"errors"
	"fmt"
	"shopping-cart/backend/models"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	discountCodePrefix           = "discount:code:"
	usedCodePrefix               = "discount:used:"
	lockPrefix                   = "discount:lock:"
	happyHoursCodesCount         = 10
	happyHoursDiscountPercentage = 0.18
	discountKey                  = "HAPPYHOURS"
	lockTimeout                  = 5 * time.Second
)

type DiscountCodeResponse struct {
	Code         string  `json:"code"`
	DiscountRate float64 `json:"discountRate"`
}

type ValidateDiscountResponse struct {
	Valid        bool    `json:"valid"`
	Code         string  `json:"code,omitempty"`
	DiscountRate float64 `json:"discountRate,omitempty"`
	Error        string  `json:"error,omitempty"`
}

type DiscountService struct {
	redisClient *redis.Client
}

func NewDiscountService(redisClient *redis.Client) *DiscountService {
	return &DiscountService{
		redisClient: redisClient,
	}
}

func (s *DiscountService) InitializeDiscountCodes(ctx context.Context) error {
	exists, err := s.redisClient.Exists(ctx, fmt.Sprintf("%s%s_%d", discountCodePrefix, discountKey, 0)).Result()
	if err != nil {
		return err
	}

	if exists > 0 {
		return nil
	}

	for i := 0; i < happyHoursCodesCount; i++ {
		key := fmt.Sprintf("%s%s_%d", discountCodePrefix, discountKey, i)
		err := s.redisClient.Set(ctx, key, "available", 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DiscountService) ValidateDiscountCode(ctx context.Context, code string) (*DiscountCodeResponse, error) {
	code = strings.ToUpper(code)

	if !strings.HasPrefix(code, discountKey) {
		return nil, errors.New(models.ErrDiscountCode)
	}

	foundAvailable := false
	var specificCode string

	for i := 0; i < happyHoursCodesCount; i++ {
		key := fmt.Sprintf("%s%s_%d", discountCodePrefix, discountKey, i)
		lockKey := fmt.Sprintf("%s%s_%d", lockPrefix, discountKey, i)

		locked, err := s.redisClient.SetNX(ctx, lockKey, "1", lockTimeout).Result()
		if err != nil {
			return nil, fmt.Errorf("failed to acquire lock: %v", err)
		}
		if !locked {
			continue
		}
		defer s.redisClient.Del(ctx, lockKey)

		txf := func(tx *redis.Tx) error {
			val, err := tx.Get(ctx, key).Result()
			if err == redis.Nil {
				return nil
			}
			if err != nil {
				return err
			}
			if val == "available" {
				specificCode = fmt.Sprintf("%s_%d", discountKey, i)
				foundAvailable = true
				_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
					pipe.Set(ctx, key, "used", 24*time.Hour)
					usedKey := fmt.Sprintf("%s%s", usedCodePrefix, specificCode)
					pipe.Set(ctx, usedKey, time.Now().String(), 24*time.Hour)
					return nil
				})
				return err
			}
			return nil
		}

		err = s.redisClient.Watch(ctx, txf, key)
		if err == nil && foundAvailable {
			break
		}
	}

	if !foundAvailable {
		return nil, errors.New("all discount codes have been used")
	}

	return &DiscountCodeResponse{
		Code:         specificCode,
		DiscountRate: happyHoursDiscountPercentage,
	}, nil
}
