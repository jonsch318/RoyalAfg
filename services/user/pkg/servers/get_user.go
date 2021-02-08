package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) GetUserById(ctx context.Context, m *protos.GetUser) (*protos.User, error) {
	s.l.Infof("Call Get By Username")

	user, err := s.db.FindById(m.GetIdentifier())

	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	msg := &protos.User{
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Id:        user.ID.Hex(),
		Username:  user.Username,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
		Email:     user.Email,
		Hash:      user.Hash,
	}

	return msg, nil
}

func (s *UserServer) GetUserByUsername(ctx context.Context, m *protos.GetUser) (*protos.User, error) {
	s.l.Infof("Call Get By Username")

	user, err := s.db.FindByUsername(m.GetIdentifier())

	if err != nil {
		s.l.Errorw("Could not find user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	msg := &protos.User{
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
		Id:        user.ID.Hex(),
		Username:  user.Username,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
		Email:     user.Email,
		Hash:      user.Hash,
	}

	return msg, nil
}
