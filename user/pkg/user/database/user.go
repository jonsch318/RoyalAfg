package database

import (
	"github.com/JohnnyS318/RoyalAfgInGo/shared/pkg/models"
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

type UserDatabase struct {
	l    *zap.SugaredLogger
	coll *mgm.Collection
}

func NewUserDatabase(logger *zap.SugaredLogger) *UserDatabase {
	coll := mgm.Coll(&models.User{})

	logger.Infof("Connected to Collection %v", coll.Name())

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

	db.coll.Create(user)
	db.logger.Info("Inserted new User ", user.Username)
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

func (db *UserDatabases) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}

	err := db.coll.FindOne(mgm.Ctx(), bson.M{"username": username}).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
