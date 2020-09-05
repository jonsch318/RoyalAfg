package servers

import (
	"context"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *UserServer) UpdateUser(ctx context.Context, m *protos.User) (*protos.User, error) {
	//user := fromMessageUser(m)

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Error(codes.Unauthenticated, "Could not read api key")
	}

	s.l.Warnf("hallo %v", md)
	return m, nil
}
