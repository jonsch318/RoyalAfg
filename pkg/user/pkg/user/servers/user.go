package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/shared/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user/pkg/user/database"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserServer struct {
	l  *zap.SugaredLogger
	db *database.UserDatabase
}

func NewUserServer(logger *zap.SugaredLogger, database *database.UserDatabase) *UserServer {
	return &UserServer{
		l:  logger,
		db: database,
	}
}

func (s *UserServer) SaveUser(ctx context.Context, m *protos.User) (*protos.User, error) {

	s.l.Infof("Called SaveUser Grpc %v", m)

	user := fromMessageUserExact(m)

	if err := user.Validate(); err != nil {
		s.l.Errorw("Validation error", "error", err)
		return nil, err
	}

	err := s.db.CreateUser(user)

	if err != nil {
		return nil, err
	}
	return toMessageUser(user), nil
}

func (s *UserServer) GetUserById(ctx context.Context, m *protos.GetUser) (*protos.User, error) {

	user, err := s.db.FindById(m.GetIdentifier())

	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	return toMessageUser(user), nil
}

func (s *UserServer) GetUserByUsername(ctx context.Context, m *protos.GetUser) (*protos.User, error) {
	user, err := s.db.FindByUsername(m.GetIdentifier())

	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	return toMessageUser(user), nil
}
func (s *UserServer) UpdateUser(ctx context.Context, m *protos.User) (*protos.User, error) {
	//user := fromMessageUser(m)

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Could not read api key")
	}

	s.l.Warnf("hallo %v", md)
	return m, nil
}

func toMessageUser(user *models.User) *protos.User {
	return &protos.User{
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Id:        user.ID.Hex(),
		Username:  user.Username,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
		Email:     user.Email,
		Hash:      user.Hash,
	}
}

func fromMessageUser(m *protos.User) *models.User {
	user := models.NewUser(m.GetUsername(), m.GetEmail(), m.GetFullName(), m.GetBirthdate())

	user.Hash = m.Hash

	return user
}

func fromMessageUserExact(m *protos.User) *models.User {
	user := models.NewUser(m.GetUsername(), m.GetEmail(), m.GetFullName(), m.GetBirthdate())

	user.Hash = m.Hash

	return user
}
