package servers

import (
	"context"

	"github.com/jonsch318/royalafg/pkg/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) GetUserById(ctx context.Context, m *protos.GetUser) (*protos.User, error) {

	//Find user
	user, err := s.db.FindById(m.GetIdentifier())
	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	//Send back user
	msg := protos.ToMessageUser(user)
	return msg, nil
}

func (s *UserServer) GetUserByUsername(ctx context.Context, m *protos.GetUser) (*protos.User, error) {

	//Find user
	user, err := s.db.FindByUsername(m.GetIdentifier())
	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	//Send back user
	msg := protos.ToMessageUser(user)
	return msg, nil
}
