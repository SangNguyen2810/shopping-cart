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
	// Redis key prefix for discount codes
	discountCodePrefix = "discount:code:"
	// Redis key prefix for used codes
	usedCodePrefix = "discount:used:"
	// Number of HAPPYHOURS codes to generate
	happyHoursCodesCount = 10
	// HAPPYHOURS discount percentage
	happyHoursDiscountPercentage = 0.18
)

// DiscountService handles discount code validation and management
type DiscountService struct {
	redisClient *redis.Client
}

// NewDiscountService creates a new discount service
func NewDiscountService(redisClient *redis.Client) *DiscountService {
	return &DiscountService{
		redisClient: redisClient,
	}
}

// InitializeDiscountCodes sets up the initial discount codes in Redis if not already present
func (s *DiscountService) InitializeDiscountCodes(ctx context.Context) error {
	// Check if codes are already initialized
	exists, err := s.redisClient.Exists(ctx, fmt.Sprintf("%s%s%d", discountCodePrefix, "HAPPYHOURS", 0)).Result()
	if err != nil {
		return err
	}

	// If codes already exist, no need to initialize
	if exists > 0 {
		return nil
	}

	// Initialize HAPPYHOURS codes (0-9)
	for i := 0; i < happyHoursCodesCount; i++ {
		key := fmt.Sprintf("%s%s%d", discountCodePrefix, "HAPPYHOURS", i)
		err := s.redisClient.Set(ctx, key, "available", 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

// ValidateDiscountCode checks if a discount code is valid and locks it if successful
func (s *DiscountService) ValidateDiscountCode(ctx context.Context, code string) (float64, error) {
	code = strings.ToUpper(code)

	if !strings.HasPrefix(code, "HAPPYHOURS") {
		return 0, errors.New("invalid discount code")
	}

	// Extract the ID from the code (should be HAPPYHOURS0 through HAPPYHOURS9)
	var codeID int
	if len(code) == 10 {
		// HAPPYHOURS (without number) - try to find an available code
		for i := 0; i < happyHoursCodesCount; i++ {
			key := fmt.Sprintf("%s%s%d", discountCodePrefix, "HAPPYHOURS", i)
			val, err := s.redisClient.Get(ctx, key).Result()
			if err == nil && val == "available" {
				codeID = i
				code = fmt.Sprintf("HAPPYHOURS%d", i)
				break
			}
		}
	} else if len(code) == 11 && code[10] >= '0' && code[10] <= '9' {
		// HAPPYHOURS with a number
		codeID = int(code[10] - '0')
	} else {
		return 0, errors.New("invalid HAPPYHOURS code format")
	}

	// Check if the code is available
	key := fmt.Sprintf("%s%s%d", discountCodePrefix, "HAPPYHOURS", codeID)
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

	// Lock the code (mark as used with a 24-hour expiration)
	usedKey := fmt.Sprintf("%s%s", usedCodePrefix, code)
	err = s.redisClient.Set(ctx, key, "used", 24*time.Hour).Err()
	if err != nil {
		return 0, err
	}

	// Record who used the code (in a real app, would store user ID)
	err = s.redisClient.Set(ctx, usedKey, time.Now().String(), 24*time.Hour).Err()
	if err != nil {
		return 0, err
	}

	// Return the discount percentage
	return happyHoursDiscountPercentage, nil
}

// GetDiscountPercentage returns the discount percentage for a given code
func (s *DiscountService) GetDiscountPercentage(code string) float64 {
	if strings.HasPrefix(strings.ToUpper(code), "HAPPYHOURS") {
		return happyHoursDiscountPercentage
	}
	return 0
}
