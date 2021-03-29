package database

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

func (db *UserDatabase) SetCache(user *models.User) error {
	return db.userCache.Set(&cache.Item{
		Ctx:            context.TODO(),
		Key:            user.ID.Hex(),
		Value:          user,
		TTL:            time.Minute * 5,
	})
}

func (db *UserDatabase) GetCache(id string) (*models.User, error) {
	user := new(models.User)
	err := db.userCache.Get(context.TODO(), id, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
