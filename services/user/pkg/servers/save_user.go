package servers

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
)

//SaveUser saves a user to the database
func (s *UserServer) SaveUser(ctx context.Context, m *protos.User) (*protos.User, error) {
	//Decode user
	user := protos.FromMessageUser(m)
	id, err := primitive.ObjectIDFromHex(m.Id)
	if err != nil {
		s.l.Errorw("Could not parse id", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user.ID = id

	//Validate User. Basically unnecessary because CreateUser already validates the user. But we can send custom error message now.
	if err = user.Validate(); err != nil {
		s.l.Errorw("Validation error", "error", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	//save user on db
	err = s.db.CreateUser(user)
	if err != nil {
		if database.IsDup(err) {
			s.l.Errorw("Username or Email duplication", "error", "Username or Email already exist")
			return nil, status.Error(codes.AlreadyExists, "Username or Email already used pleace try again using a different Username or Email")
		}

		s.l.Errorw("Error during parsing", "error", err)
		return nil, status.Error(codes.Internal, "User could not be saved to the database")
	}

	//Metrics
	s.metrics.SavedUser()

	return protos.ToMessageUser(user), nil
}
