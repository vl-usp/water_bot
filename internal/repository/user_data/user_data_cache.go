package user_data

import (
	"context"
	"strconv"
	"time"

	"github.com/vl-usp/water_bot/internal/model"
	"github.com/vl-usp/water_bot/internal/repository/user_data/converter"
)

const (
	cachePrefix = "user_data_"
	cacheTTL    = 120 * time.Minute
)

// SetToCache sets user data to cache.
func (r *repo) SaveField(ctx context.Context, userID int64, field string, value interface{}) error {
	userIDString := strconv.Itoa(int(userID))
	err := r.cache.HSetField(ctx, cachePrefix+userIDString, field, value)
	if err != nil {
		return err
	}

	// logger.Get("repository", "r.SaveField").Debug("field saved", "userID", userID, "field", field, "value", value)

	return r.cache.Expire(ctx, cachePrefix+userIDString, cacheTTL)
}

// GetFromCache returns user data from cache.
func (r *repo) GetFromCache(ctx context.Context, userID int64) (*model.UserData, error) {
	data, err := r.cache.HGetAll(ctx, cachePrefix+strconv.Itoa(int(userID)))
	if err != nil {
		return nil, err
	}

	userData, err := converter.ToUserDataFromCache(data)
	if err != nil {
		return nil, err
	}

	userData.UserID = userID

	// logger.Get("repository", "r.GetFromCache").Debug("user data from cache", "userID", userID, "userData", userData)

	return userData, nil
}
