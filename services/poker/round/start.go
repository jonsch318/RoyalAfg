package round

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/random"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

func (r *Round) Start(players []models.Player, publicPlayers []models.PublicPlayer, dealer int) {
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v", r)
			debug.PrintStack()
		}
	}()
	r.Players = players
	r.PublicPlayers = publicPlayers

	sleepTime := viper.GetDuration(serviceconfig.StepSleepDuration)

	log.Logger.Debugf("Start called reseting bank and initializing")

	//Initializing start
	r.Bank.Reset()
	r.Dealer = dealer
	r.Players = players
	r.InCount = byte(len(players))
	r.HoleCards = make(map[string][2]models.Card, len(players))
	log.Logger.Debugf("Configured round")

	//We let the client handle everything and then start the round.
	time.Sleep(sleepTime)

	//publish players and position
	log.Logger.Debugf("Publishing players")
	r.Bank.UpdatePublicPlayerBuyIn(r.PublicPlayers)
	log.Logger.Debugf("Transmiting Public Player info %v", r.PublicPlayers)
	for i := range r.Players {
		if r.Players[i].ID != r.PublicPlayers[i].ID {
			log.Logger.Errorf("Public-Private Player Information unsynchronized %v", r.Players[i].Username)
		}
		if err := utils.SendToPlayerInListTimeout(r.Players, i, events.NewGameStartEvent(r.PublicPlayers, i, r.Bank.Pot.Display())); err != nil {
			log.Logger.Debugf("Error during game start event transmittion %v", err.Error())
			_ = r.Leave(r.Players[i].ID)
		}
	}

	r.WhileNotEnded(func(){
		r.sendDealer()
	})
	// Publish chosen Dealer
	time.Sleep(sleepTime)

	r.WhileNotEnded(func(){
		//set predefined blinds
		err := r.setBlinds()
		if err != nil {
			r.Bank.ConcludeRound(nil, r.PublicPlayers)
			return
		}
	})

	cards, err := random.SelectCards(5 + 2*int(r.InCount))

	if err != nil {
		log.Logger.Errorw("error during card generation")
		return
	}

	log.Logger.Debugf("Cards generated %v", cards)

	// Fill the board
	for i := 0; i < 5; i++ {
		r.Board[i] = cards[i]
	}

	// Set players hole cards
	r.holeCards(cards[4:])
	log.Logger.Infof("Hole cards set")

	time.Sleep(sleepTime)

	r.WhileNotEnded(func() {
		r.actions(true)
	})
	log.Logger.Debugf("Generate cards")

	for i := 3; i < 6; i++ {
		r.WhileNotEnded(func() {
			log.Logger.Debugf("Started action round [%v]", i-2)
			//Send the board cards first 3 then the 4th and then the 5th
			r.SendBoardEvent(i)

			//Acquire the actions of the players
			r.actions(false)
		})
	}

	//Done

	//Evaluation
	r.Evaluate()
	time.Sleep(sleepTime)
}

//Evaluate concludes this round and publishes all results to the bank service for performing the real transactions.
func (r *Round) Evaluate() {
	//Determine winner(s) of this round. Most of the time one but can be more if exactly equal cards.
	winners := showdown.Evaluate(r.Players, r.HoleCards, r.Board, r.InCount)
	log.Logger.Infow("Winners determined: %v", winners)
	r.LogCards()

	//Publish commands to bank service.
	shares := r.Bank.ConcludeRound(winners, r.PublicPlayers)


	//Send winning results to clients. You could add the hole cards for clarity. But this can be added fairly easily.
	winningPublic := make([]models.PublicPlayer, len(winners))
	for i, w := range winners {
		if r.PublicPlayers[w.Position].ID != w.Player.ID {
			log.Logger.Errorf("Player and win info not synchronized")
		}
		winningPublic[i] = r.PublicPlayers[w.Position]
		log.Logger.Debugf("Winning public %v", winningPublic[i])
	}
	log.Logger.Debugf("Winning Publics %v", winningPublic)
	utils.SendToAll(r.Players, events.NewGameEndEvent(winningPublic, shares[0]))
}

func (r *Round) LogCards() {
	for _, player := range r.Players {
		str := fmt.Sprintf("%s Cards: [ ", player.Username)
		for _, card := range r.Board {
			str += card.String() + ", "
		}
		card0 := r.HoleCards[player.ID][0]
		card1 := r.HoleCards[player.ID][1]

		str += card0.String() + ", "
		str += card1.String() + " ]"

		log.Logger.Debugf(str)
	}
}

//SendBoardEvent is a little utility for sorting the right board event name for a given number of cards
func (r *Round) SendBoardEvent(cardCount int) {
	switch cardCount {
	case 3:
		utils.SendToAll(r.Players, events.NewFlopEvent(r.Board))
	case 4:
		utils.SendToAll(r.Players, events.NewTurnEvent(r.Board))

	case 5:
		utils.SendToAll(r.Players, events.NewRiverEvent(r.Board))

	default:
		log.Logger.Errorf("SendBoardEvent with cardCount not between 3-5: %v", cardCount)
	}
}

func (r *Round) End() {
	log.Logger.Error("ending round due to error")
	r.Ended = true
}
