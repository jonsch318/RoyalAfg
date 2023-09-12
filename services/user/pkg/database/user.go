package database

import (
	"github.com/go-redis/cache/v8"

	"github.com/jonsch318/royalafg/pkg/models"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoUserDB struct {
	l         *zap.SugaredLogger
	userCache *cache.Cache
	coll      *mgm.Collection
}

func NewUserDatabase(logger *zap.SugaredLogger, userCache *cache.Cache) *MongoUserDB {
	coll := mgm.Coll(&models.User{})

	logger.Infof("Connected to Collection %v", coll.Name())

	i := []mongo.IndexModel{
		{
			Keys:    bson.M{"username": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"email": 1},
			Options: options.Index().SetUnique(true),
		},
	}

	ind, err := coll.Indexes().CreateMany(
		mgm.Ctx(),
		i,
	)

	for _, in := range ind {
		logger.Debugf("Index created %v", in)
	}

	if err != nil {
		logger.Errorw("Error during creating indexes", "error", err)
		return nil
	}

	return &MongoUserDB{
		l:         logger,
		userCache: userCache,
		coll:      coll,
	}
}
func (db *MongoUserDB) CreateUser(user *models.User) error {

	err := user.Validate()

	if err != nil {
		return err
	}

	db.l.Info("Succeeded validation")

	err = db.coll.Create(user)

	db.l.Infof("User Creation %v", err)

	if err != nil {
		return err
	}

	db.l.Info("Inserted new User ", user.Username)
	err = db.SetCache(user)
	if err != nil {
		db.l.Debugf("Could not set cache %v", err)
	}
	return nil
}

func (db *MongoUserDB) UpdateUser(user *models.User) error {
	db.l.Debugf("Succeeded validation")

	err := db.coll.Update(user)

	db.l.Infof("User Updated %v", err)

	if err != nil {
		db.l.Infof("Error during user update %v", err)
		return err
	}
	db.l.Infof("Updated user %v", user.GetID())

	err = db.SetCache(user)
	if err != nil {
		db.l.Debugf("Could not set cache %v", err)
	}
	return nil
}

func (db *MongoUserDB) DeleteUser(user *models.User) error {
	return db.coll.Delete(user)
}

// FindById returns the user, if found, with the given id. This is cached
func (db *MongoUserDB) FindById(id string) (*models.User, error) {

	//check cache of id
	user, cacheHit, err := db.GetCache(id)
	if err != nil {
		db.l.Debugf("Could not get cache although cache hit %v", err)
	}

	if user != nil && cacheHit {
		return user, nil
	}

	user = &models.User{}
	err = db.coll.FindByID(id, user)

	if err != nil {
		return nil, err
	}

	err = db.SetCache(user)
	if err != nil {
		db.l.Debugf("Could not set cache %v", err)
	}

	return user, nil
}

// FindByEmail returns the user, if found, with the given email. This is NOT cached
func (db *MongoUserDB) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, err
	}

	err = db.SetCache(user)
	if err != nil {
		db.l.Debugf("Could not set cache %v", err)
	}

	return user, nil
}

// FindByUsername returns the user, if found, with the given username. This is NOT cached
func (db *MongoUserDB) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(user)

	if err != nil {
		return nil, err
	}

	err = db.SetCache(user)
	if err != nil {
		db.l.Debugf("Could not set cache %v", err)
	}

	return user, nil
}
