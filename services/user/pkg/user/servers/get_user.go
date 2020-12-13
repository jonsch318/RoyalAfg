package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/services/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
