package utils

import (
	"time"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func DisconnectClient(logger *zap.SugaredLogger, client *mongo.Client) {
	err := client.Disconnect(mgm.Ctx())
	if err != nil {
		logger.Fatalw("MongoDB could not disconnect", "error", err)
	}
	logger.Warn("Mongodb disconnected")
}

func StartMongoServerWithFatal(logger *zap.SugaredLogger, url, databaseName string, timeout time.Duration) (*mgm.Config, *mongo.Client, *mongo.Database) {
	cfg, client, db, err := StartMongoServer(url, databaseName, timeout)

	if err != nil {
		logger.Fatalw("Could not connect to the mongo database", "error", err)
	}

	logger.Warnf("Connected to Mongodb database %v", databaseName)

	return cfg, client, db
}

func StartMongoServer(url, databaseName string, timeout time.Duration) (*mgm.Config, *mongo.Client, *mongo.Database, error) {
	cfg := &mgm.Config{CtxTimeout: timeout}
	err := mgm.SetDefaultConfig(cfg, databaseName, options.Client().ApplyURI(url))

	if err != nil {
		return nil, nil, nil, err
	}

	return mgm.DefaultConfigs()

}
