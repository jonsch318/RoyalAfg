package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserStatusServer is a grpc server handler to get the status of a user (online, offline, ...)
func (s *UserServer) GetUserStatus(ctx context.Context, m *protos.UserStatusRequest) (*protos.UserStatusResponse, error) {
	//Get online status from redis as a routine
	statusChan := make(chan *database.OnlineStatus)
	go func() {

		status, err := s.statusDb.GetOnlineStatus(m.GetId())
		if err != nil {
			s.l.Errorw("Could not get online status", "error", err)
			statusChan <- nil
		}
		statusChan <- status
	}()

	//Get banned status from database

	user, err := s.db.FindById(m.GetId())

	if err != nil {
		s.l.Errorw("Could not query user", "error", err)
		return nil, status.Error(codes.NotFound, "The user could not be found")
	}

	status := <-statusChan

	protoStatus := &protos.UserStatusResponse_OnlineStatus{
		Status: int32(status.Status),
		GameId: status.GameId,
	}

	resp := &protos.UserStatusResponse{
		Banned:   uint32(user.Banned),
		Verified: uint32(user.Verified),
		Status:   protoStatus,
	}
	return resp, nil
}
