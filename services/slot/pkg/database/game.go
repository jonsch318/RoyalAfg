package database

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type GameDatabase struct {
	l    *zap.SugaredLogger
	coll *mgm.Collection
}

func NewGameDatabase(logger *zap.SugaredLogger) (*GameDatabase, error) {
	coll := mgm.Coll(&models.SlotGame{})

	logger.Infof("Created internal Collection %v", coll.Name())

	i := []mongo.IndexModel{
		{
			Keys:    bson.M{"alpha": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"beta": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	indexes, err := coll.Indexes().CreateMany(
		mgm.Ctx(),
		i,
	)

	if err != nil {
		logger.Errorw("Error during creating indexes", "error", err)
		return nil, err
	}

	for _, index := range indexes {
		logger.Debugf("Index created %v", index)
	}

	return &GameDatabase{
		l:    logger,
		coll: coll,
	}, nil
}

func (db *GameDatabase) CheckId(id string) bool {
	res := db.coll.FindOne(mgm.Ctx(), bson.M{"id": id})
	return res.Err() == nil
}

func (db *GameDatabase) SaveGameBuffer(games []*models.SlotGame) error {

	db.l.Info("Saving Buffered Games  ")

	// Create a slice of interfaces
	gamesCopies := make([]interface{}, len(games))
	for i, game := range games {
		gamesCopies[i] = *game
	}

	res, err := db.coll.InsertMany(mgm.Ctx(), gamesCopies)

	if err != nil {
		db.l.Errorw("Error during creating new game", "error", err)
		return err
	}

	db.l.Infof("Game created with id %v", res.InsertedIDs...)

	return nil
}
