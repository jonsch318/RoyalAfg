package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/dtos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/mw"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/responses"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/user"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/utils"
	"github.com/JohnnyS318/RoyalAfgInGo/services/slot/pkg/logic"
)

const SlotGameName = "Slot"

type SpinRequestDTO struct {
	factor uint
}

// POST request
// Spin: draws a new number and returns the result and prove.
func (s *SlotServer) Spin(rw http.ResponseWriter, r *http.Request) {

	// CSRF Validattion

	if err := mw.ValidateCSRF(r); err != nil {
		s.l.Errorw("Could not validate csrf token.", "error", err)
		return
	}

	// Check Permissions and login
	claims := mw.FromUserTokenContext(r.Context().Value("user"))

	eligeable := make(chan bool)

	ctx, _ := context.WithTimeout(r.Context(), time.Millisecond*200)

	// Read request body
	dto := &SpinRequestDTO{}
	err := utils.FromJSON(dto, r.Body)

	if err != nil {
		s.l.Errorw("Could not parse request body.", "error", err)
		responses.Error(rw, "Could not parse request body", http.StatusBadRequest)
		return
	}

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

	//generate the game not yet save it
	gameChan := make(chan *models.SlotGame)
	go func() {
		game, err := s.gameProvider.NewGame(dto.factor)

		if err != nil {
			gameChan <- nil
			return
		}

		gameChan <- game
	}()

	//wait for the game and the user status
	if eligible := <-eligeable; !eligible {
		s.l.Errorw("User is not eligible to play.")
		responses.Error(rw, "User is not eligible to play.", http.StatusForbidden)
		return
	}
	game := <-gameChan

	if game == nil {
		s.l.Errorw("Could not generate game.")
		responses.Error(rw, "Could not generate game.", http.StatusInternalServerError)
		return
	}

	// TODO: Connect to Bank and transact 1€
	winResult, dir := logic.ProcessGame(game, dto.factor)
	bankCommand := &bank.Command{
		UserId: claims.ID,
		Amount: &dtos.CurrencyDto{
			Value:    winResult.Amount(),
			Currency: winResult.Currency().Code,
		},
		Time:  time.UnixMicro(game.Time),
		Game:  SlotGameName,
		Lobby: fmt.Sprint(game.Time),
	}
	if dir {
		bankCommand.CommandType = bank.Withdraw
	} else {
		bankCommand.CommandType = bank.Deposit
	}

	err = s.bankService.PublishCommand(bankCommand)

	if err != nil {
		s.l.Errorw("Could not communicate with bank.", "error", err)
		responses.Error(rw, "Could not communicate with bank", http.StatusInternalServerError)
		return
	}

	//save the game
	err = s.gameProvider.SaveGame(game)

	if err != nil {

		bankCommand.CommandType = bank.Rollback

		s.l.Errorw("Could not save game.", "error", err)
		responses.Error(rw, "Could not save game.", http.StatusInternalServerError)
		return
	}

	s.l.Debugf("Game saved: %v", game)

	// TODO: Draw a new number

	// TODO: Calculate win

	// TODO: Save the result in the buffer/db

	// TODO: Connect to Bank and transact the win

	// TODO: Return the result and prove
}
