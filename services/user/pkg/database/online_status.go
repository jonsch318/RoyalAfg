package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type OnlineStatus struct {
	Status byte   `json:"status"`
	GameId string `json:"gameId"`
}

type OnlineStatusDatabase struct {
	l      *zap.SugaredLogger
	client *redis.Client
}

func NewOnlineStatusDatabase(logger *zap.SugaredLogger, client *redis.Client) *OnlineStatusDatabase {
	return &OnlineStatusDatabase{
		l:      logger,
		client: client,
	}
}

func (db *OnlineStatusDatabase) SetOnlineStatus(id string, status *OnlineStatus) error {

	raw, _ := json.Marshal(status)

	return db.client.Set(context.TODO(), id, raw, 2*time.Hour).Err()
}

func (db *OnlineStatusDatabase) GetOnlineStatus(id string) (*OnlineStatus, error) {

	status := new(OnlineStatus)
	raw, err := db.client.Get(context.TODO(), id).Result()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(raw), status)

	return status, err
}
