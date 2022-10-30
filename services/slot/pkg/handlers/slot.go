package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/logic"
	"go.uber.org/zap"
)

type SlotServer struct {
	l            *zap.SugaredLogger
	gameProvider *logic.GameProvider
	userService  protos.UserServiceClient
}

func NewSlotServer(logger *zap.SugaredLogger, gameProvider *logic.GameProvider, userService protos.UserServiceClient) *SlotServer {
	return &SlotServer{
		l:            logger,
		gameProvider: gameProvider,
		userService:  userService,
	}
}
