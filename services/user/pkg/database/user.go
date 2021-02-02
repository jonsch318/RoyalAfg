package database

import (
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type UserDatabase struct {
	l    *zap.SugaredLogger
	coll *mgm.Collection
}

func NewUserDatabase(logger *zap.SugaredLogger) *UserDatabase {
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

	return &UserDatabase{
		l:    logger,
		coll: coll,
	}
}
func (db *UserDatabase) CreateUser(user *models.User) error {

	err := user.Validate()

	if err != nil {
		return err
	}

	db.l.Info("Succeded validation")

	err = db.coll.Create(user)

	db.l.Infof("User Creation %v", err)

	if err != nil {
		return err
	}

	db.l.Info("Inserted new User ", user.Username)
	return nil
}

func (db *UserDatabase) DeleteUser(user *models.User) error {
	return db.coll.Delete(user)
}

func (db *UserDatabase) FindById(id string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindByID(id, user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *UserDatabase) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *UserDatabase) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// IsDup returns whether err informs of a duplicate key error because
// a primary key index or a secondary unique index already has an entry
// with the given value.
func IsDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
