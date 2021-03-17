package servers

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
)



func (s *UserServer) UpdateUser(ctx context.Context, m *protos.User) (*protos.User, error) {

	user, err := s.db.FindById(m.GetId())

	if err != nil {
		s.l.Errorw("User query", "error", err)
		return nil, status.Error(codes.NotFound, "user with id could not be found")
	}

	if err := validation.Validate(m.Email, validation.Required, is.Email); err == nil {
		user.Email = m.Email
	}

	if err := validation.Validate(m.FullName, validation.Required); err == nil {
		user.FullName = m.FullName
	}

	err = s.db.UpdateUser(user)

	if err != nil {
		s.l.Errorw("Error during user update", "error", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	return toMessageUser(user), nil
}
