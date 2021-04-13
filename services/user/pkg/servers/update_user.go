package servers

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
)


//Update user updates a user with the given id.
func (s *UserServer) UpdateUser(ctx context.Context, m *protos.User) (*protos.User, error) {

	//Find user
	user, err := s.db.FindByUsername(m.GetId())
	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	//Check if email changed
	if err = validation.Validate(m.Email, validation.Required, is.Email); err == nil {
		user.Email = m.Email
	}

	//Check if fullName changed
	if err = validation.Validate(m.FullName, validation.Required); err == nil {
		user.FullName = m.FullName
	}

	//Update user on db
	err = s.db.UpdateUser(user)
	if err != nil {
		s.l.Errorw("Error during user update", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}


	return protos.ToMessageUser(user), nil
}
