package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	discountCodePrefix           = "discount:code:"
	usedCodePrefix               = "discount:used:"
	happyHoursCodesCount         = 10
	happyHoursDiscountPercentage = 0.18
	discountKey                  = "HAPPYHOURS"
)

type DiscountService struct {
	redisClient *redis.Client
}

func NewDiscountService(redisClient *redis.Client) *DiscountService {
	return &DiscountService{
		redisClient: redisClient,
	}
}

func (s *DiscountService) InitializeDiscountCodes(ctx context.Context) error {
	exists, err := s.redisClient.Exists(ctx, fmt.Sprintf("%s%s%d", discountCodePrefix, discountKey, 0)).Result()
	if err != nil {
		return err
	}

	if exists > 0 {
		return nil
	}

	for i := 0; i < happyHoursCodesCount; i++ {
		key := fmt.Sprintf("%s%s%d", discountCodePrefix, discountKey, i)
		err := s.redisClient.Set(ctx, key, "available", 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DiscountService) ValidateDiscountCode(ctx context.Context, code string) (float64, error) {
	code = strings.ToUpper(code)

	if !strings.HasPrefix(code, "HAPPYHOURS") {
		return 0, errors.New("invalid discount code")
	}

	var codeID int
	if len(code) == 10 {
		foundAvailable := false
		for i := 0; i < happyHoursCodesCount; i++ {
			key := fmt.Sprintf("%s%s%d", discountCodePrefix, discountKey, i)
			val, err := s.redisClient.Get(ctx, key).Result()
			if err == nil && val == "available" {
				codeID = i
				code = fmt.Sprintf("HAPPYHOURS%d", i)
				foundAvailable = true
				break
			}
		}
		if !foundAvailable {
			return 0, errors.New("all discount codes have been used")
		}
	} else if len(code) == 11 && code[10] >= '0' && code[10] <= '9' {
		codeID = int(code[10] - '0')
	} else {
		return 0, errors.New("invalid HAPPYHOURS code format")
	}

	key := fmt.Sprintf("%s%s%d", discountCodePrefix, discountKey, codeID)
	val, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.New("discount code not found")
		}
		return 0, err
	}

	if val != "available" {
		return 0, errors.New("discount code has already been used")
	}

	usedKey := fmt.Sprintf("%s%s", usedCodePrefix, code)
	err = s.redisClient.Set(ctx, key, "used", 24*time.Hour).Err()
	if err != nil {
		return 0, err
	}

	err = s.redisClient.Set(ctx, usedKey, time.Now().String(), 24*time.Hour).Err()
	if err != nil {
		return 0, err
	}

	return happyHoursDiscountPercentage, nil
}

func (s *DiscountService) GetDiscountPercentage(code string) float64 {
	if strings.HasPrefix(strings.ToUpper(code), discountKey) {
		return happyHoursDiscountPercentage
	}
	return 0
}
