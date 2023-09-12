package handlers

import (
	"github.com/jonsch318/royalafg/pkg/bank"
	"github.com/jonsch318/royalafg/pkg/protos"
	"github.com/jonsch318/royalafg/services/slot/pkg/logic"
	"go.uber.org/zap"
)

type SlotServer struct {
	l            *zap.SugaredLogger
	gameProvider *logic.GameProvider
	userService  protos.UserServiceClient
	bankService  *bank.RabbitBankConnection
}

func NewSlotServer(logger *zap.SugaredLogger, gameProvider *logic.GameProvider, userService protos.UserServiceClient, bankService *bank.RabbitBankConnection) *SlotServer {
	return &SlotServer{
		l:            logger,
		gameProvider: gameProvider,
		userService:  userService,
		bankService:  bankService,
	}
}
