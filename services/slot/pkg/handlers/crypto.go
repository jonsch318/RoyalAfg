package handlers

import (
	"net/http"

	"github.com/jonsch318/royalafg/pkg/dtos"
	"github.com/jonsch318/royalafg/pkg/utils"
	"github.com/jonsch318/royalafg/services/slot/pkg/crypto"
)

func (s *SlotServer) CryptoInfo(rw http.ResponseWriter, r *http.Request) {
	pkKey := s.gameProvider.GetPublicKey()
	pkEncoded, err := crypto.ExportPublicKey(pkKey)

	if err != nil {
		err = utils.RespondWithError(rw, http.StatusInternalServerError, err.Error())
		if err != nil {
			s.l.Errorw("Failed to write response", "err", err)
		}
		return
	}

	err = utils.ToJSON(&dtos.CryptoInfoDto{PublicKey: pkEncoded}, rw)

	if err != nil {
		s.l.Errorw("Failed to write response", "err", err)
	}
}
