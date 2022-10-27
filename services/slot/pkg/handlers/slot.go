package handlers

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/crypto"
	"go.uber.org/zap"
)

type SlotServer struct {
	l   *zap.SugaredLogger
	rng *crypto.VRFNumberGenerator
}

func NewSlotServer(logger *zap.SugaredLogger, rng *crypto.VRFNumberGenerator) *SlotServer {
	return &SlotServer{
		l:   logger,
		rng: rng,
	}
}
