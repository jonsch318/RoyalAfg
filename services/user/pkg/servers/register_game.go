package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *UserServer) RegisterGame(ctx context.Context, m *protos.RegisterGameRequest) (*protos.OnlineStatus, error) {
	//Decode user

	//Validate User? Currently the grpc service assumes only internat services can call it.

	user, err := s.db.FindById(m.Id)

	if err != nil {
		return nil, status.Error(codes.NotFound, "The user with the id could not be found")
	}

	state := database.OnlineStatus{
		Status: database.Online,
		GameId: m.GameId,
	}

	err = s.statusDb.SetOnlineStatus(user.ID.String(), &state)

	return &protos.OnlineStatus{
		Status: int32(state.Status),
		GameId: state.GameId,
	}, err
}
