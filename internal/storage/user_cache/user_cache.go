package user_cache

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/storage/user_cache/converter"
)

const (
	cachePrefix = "user_params_"
	cacheTTL    = 120 * time.Minute
)

// SetToCache sets user data to cache.
func (s *store) SaveUserParam(ctx context.Context, userID int64, field string, value interface{}) error {
	userIDString := strconv.Itoa(int(userID))
	err := s.cache.HSetField(ctx, cachePrefix+userIDString, field, value)
	if err != nil {
		return fmt.Errorf("failed to save field to cache: %w", err)
	}

	// logger.Get("repository", "r.SaveField").Debug("field saved", "userID", userID, "field", field, "value", value)

	return s.cache.Expire(ctx, cachePrefix+userIDString, cacheTTL)
}

// GetFromCache returns user data from cache.
func (s *store) GetUserParams(ctx context.Context, userID int64) (*model.UserParams, error) {
	data, err := s.cache.HGetAll(ctx, cachePrefix+strconv.Itoa(int(userID)))
	if err != nil {
		return nil, fmt.Errorf("failed to get user data from cache: %w", err)
	}

	// logger.Get("repository", "r.GetFromCache").Debug("user data from cache", "userID", userID, "data", data)

	userData, err := converter.ToUserParamsFromCache(data)
	if err != nil {
		return nil, fmt.Errorf("failed to convert user data from cache: %w", err)
	}

	// logger.Get("repository", "r.GetFromCache").Debug("converted user data from cache", "userID", userID, "userData", userData)

	return userData, nil
}
