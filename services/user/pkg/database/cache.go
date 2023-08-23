package database

import (
	"context"
	"time"

	"github.com/go-redis/cache/v8"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

func (db *MongoUserDB) SetCache(user *models.User) error {
	return db.userCache.Set(&cache.Item{
		Ctx:   context.TODO(),
		Key:   user.ID.Hex(),
		Value: user,
		TTL:   time.Minute * 10,
	})
}

func (db *MongoUserDB) GetCache(id string) (*models.User, bool, error) {
	user := new(models.User)
	if !db.userCache.Exists(context.TODO(), id) {
		return nil, false, nil
	}
	err := db.userCache.Get(context.TODO(), id, user)
	if err != nil {
		return nil, true, err
	}
	return user, true, nil
}
