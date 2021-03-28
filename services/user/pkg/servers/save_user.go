package servers

import (
	"context"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) SaveUser(ctx context.Context, m *protos.User) (*protos.User, error) {

	s.l.Infof("Called SaveUser Grpc %v", m)

	user := &models.User{
		Username:  m.Username,
		Email:     m.Email,
		Birthdate: m.Birthdate,
		FullName:  m.FullName,
		Hash:      m.Hash,
	}

	id, err := primitive.ObjectIDFromHex(m.Id)

	if err != nil {
		s.l.Errorw("Could not parse id", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	user.ID = id

	user.CreatedAt = time.Unix(m.CreatedAt, 0)
	user.UpdatedAt = time.Unix(m.UpdatedAt, 0)

	if err := user.Validate(); err != nil {
		s.l.Errorw("Validation error", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Validate unique Username and Email

	err = s.db.CreateUser(user)

	if err != nil {

		if database.IsDup(err) {
			s.l.Errorw("Username or Email duplication", "error", "Username or Email already exist")
			return nil, status.Error(codes.AlreadyExists, "Username or Email already used pleace try again using a different Username or Email")
		}

		s.l.Errorw("Error during parsing", "error", err)
		return nil, status.Error(codes.Internal, "User could not be saved to the database")
	}

	s.metrics.SavedUser()

	return toMessageUser(user), nil
}
