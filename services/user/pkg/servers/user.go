package servers

import (
	"go.uber.org/zap"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/database"
	"github.com/JohnnyS318/RoyalAfgInGo/services/user/pkg/metrics"
)

// UserServer is a grpc server handler to save, update or retrieve a user from the database
type UserServer struct {
	protos.UnimplementedUserServiceServer
	l        *zap.SugaredLogger
	db       database.UserDB
	statusDb database.OnlineStatusDB
	metrics  *metrics.User
}

// NewUserServer create a new grpc user server
func NewUserServer(logger *zap.SugaredLogger, database database.UserDB, statusDb database.OnlineStatusDB, metrics *metrics.User) *UserServer {
	return &UserServer{
		l:        logger,
		db:       database,
		statusDb: statusDb,
		metrics:  metrics,
	}
}
