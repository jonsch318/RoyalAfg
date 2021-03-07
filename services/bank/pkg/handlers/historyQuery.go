package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/services/bank/pkg/dtos"
)

type HistoryQueryDto struct {
	UserID  string                     `json:"userId"`
	History []dtos.AccountHistoryEvent `json:"history"`
}

func (h *Account) QueryHistory(rw http.ResponseWriter, r *http.Request) {
	//vals := r.URL.Query()
	//i, _ := strconv.Atoi(vals["i"][0])

	i := 0

	claims := mw.FromUserTokenContext(r.Context().Value("user"))
	log.Logger.Debugw("Claims decoded", claims)
	history, err := h.historyReadModel.GetAccountHistory(claims.ID, i, 25)

	if err != nil {
		log.Logger.Errorw("Query error", "error", err, "ID", claims.ID)
		http.Error(rw, err.Error(), http.StatusNotFound)
		return
	}

	log.Logger.Debugw("History Query sending", "history", history)

	_ = json.NewEncoder(rw).Encode(&HistoryQueryDto{
		UserID:  claims.ID,
		History: history,
	})

}
