package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user"
)

// POST request
// Spin: draws a new number and returns the result and prove.
func (s *SlotServer) Spin(rw http.ResponseWriter, r *http.Request) {

	// TODO: CSRF Validattion

	if err := mw.ValidateCSRF(r); err != nil {
		s.l.Errorw("Could not validate csrf token.", "error", err)
		return
	}

	// TODO: Check Permissions and login
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	eligeable := make(chan bool)

	ctx, _ := context.WithTimeout(r.Context(), time.Millisecond*200)

	go func() {

		resp, err := s.userService.GetUserStatus(ctx, &protos.UserStatusRequest{
			Id: claims.ID,
		})

		if err != nil {
			eligeable <- false
			return
		}
		eligeable <- user.IsPlayEligible(uint8(resp.GetVerified()), uint8(resp.GetBanned()))

	}()

	gameChan := make(chan *models.SlotGame)

	go func() {
		game, err := s.gameProvider.NewGame()

		if err != nil {
			gameChan <- nil
			return
		}

		gameChan <- game
	}()

	// TODO: Read request body

	// TODO: Connect to Bank and transact 1€

	// TODO: Draw a new number

	// TODO: Calculate win

	// TODO: Save the result in the buffer/db

	// TODO: Connect to Bank and transact the win

	// TODO: Return the result and prove
}
