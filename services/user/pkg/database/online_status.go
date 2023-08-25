package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const Online = 1
const Offline = 0

type OnlineStatus struct {
	Status byte   `json:"status"`
	GameId string `json:"gameId"`
}

type RedisStatusDatabase struct {
	l      *zap.SugaredLogger
	client *redis.Client
}

func NewOnlineStatusDatabase(logger *zap.SugaredLogger, client *redis.Client) *RedisStatusDatabase {
	return &RedisStatusDatabase{
		l:      logger,
		client: client,
	}
}

func (db *RedisStatusDatabase) SetOnlineStatus(id string, status *OnlineStatus) error {

	raw, err := json.Marshal(status)

	if err != nil {
		return err
	}

	return db.client.Set(context.TODO(), id, raw, 2*time.Hour).Err()
}

func (db *RedisStatusDatabase) GetOnlineStatus(id string) (*OnlineStatus, error) {

	status := new(OnlineStatus)
	raw, err := db.client.Get(context.TODO(), id).Result()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(raw), status)

	return status, err
}
