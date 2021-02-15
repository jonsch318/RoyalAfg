package round

import (
	"time"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/showdown"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

func (r *Round) Start(players []models.Player, publicPlayers []models.PublicPlayer, dealer int) {
	defer func(){
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v", r)
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
	for i := range r.Players {
		utils.SendToPlayerInList(r.Players, i, events.NewGameStartEvent(publicPlayers, i))
	}

	// Publish chosen Dealer
	r.sendDealer()
	time.Sleep(sleepTime)

	//set predefined blinds
	err := r.setBlinds()
	if err != nil {
		r.Bank.ConcludeRound(nil)
		return
	}

	// Set players hole cards
	holeCards(r.Players, r.HoleCards, r.cardGen)
	log.Logger.Infof("Hole cards set")


	time.Sleep(sleepTime)

	r.WhileNotEnded(func(){
		r.actions(true)
	})
	log.Logger.Debugf("Generate cards")

	r.WhileNotEnded(func() {
		// Flop, turn and river cards are generated here. Doing it once lets us optimize the randomization of it.
		for i := 0; i < 5; i++ {
			r.Board[i] = r.cardGen.SelectRandom()
		}
	})

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
	winners := showdown.Evaluate(r.Players, r.HoleCards, r.Board)
	log.Logger.Infow("Winners determined", "winners", winners)

	//Publish commands to bank service.
	shares := r.Bank.ConcludeRound(winners)
	r.Bank.UpdatePublicPlayerBuyIn(r.PublicPlayers)

	//Send winning results to clients. You could add the hole cards for clarity. But this can be added fairly easily.
	winningPublic := make([]models.PublicPlayer, len(winners))
	for _, w := range winners {
		winningPublic = append(winningPublic, r.PublicPlayers[w.Position])
	}
	utils.SendToAll(r.Players, events.NewGameEndEvent(winningPublic, shares[0]))
}


//SendBoardEvent is a little utility for sorting the right board event name for a given number of cards
func (r *Round) SendBoardEvent(cardCount int){
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
