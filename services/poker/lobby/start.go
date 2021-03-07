package lobby

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"

	"github.com/spf13/viper"
)

//Start  starts a poker game
func (l *Lobby) Start() {
	// SETUP
	go func() {
		for l.Count() >= viper.GetInt(serviceconfig.PlayersRequiredForStart) {
			log.Logger.Debug("Start timer")

			// Protection against multiple games using a buffered channel. Once one is through the 15 seconds timeout all other cancel the starting process.
			timer := time.NewTimer(time.Duration(viper.GetInt(serviceconfig.GameStartTimeout)) * time.Second)
			select {
			case <-l.c:
				// game has already been called this instance is unnecessary
				log.Logger.Debugf("Game has already begun.")
				return
			case <-timer.C:
			}

			// channel is empty, so the buffer is free to be filled.
			l.c <- true

			log.Logger.Debugf("Preparing for game")

			l.PrepareForRound()

			log.Logger.Debugf("Game setup playercount: %d", len(l.Players))

			if viper.GetBool(serviceconfig.NeedEnterToStart) {
				log.Logger.Warnf("Needs input to start")
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("ENTER for game start \n")
				_, _, _ = reader.ReadLine()
			}

			if len(l.Players) < viper.GetInt(serviceconfig.PlayersRequiredForStart) {
				// Not enough players to start
				log.Logger.Errorf("Not enough players to continue")
				return
			}

			if l.dealer < 0 {
				rand.Seed(time.Now().UnixNano())
				l.dealer = rand.Intn(len(l.Players))
			} else {
				l.dealer = (l.dealer + 1) % len(l.Players)
			}

			log.Logger.Debugf("Dealer chosen %v", l.dealer)

			for i := range l.Players {
				// Set player states to active
				l.Players[i].Active = true
			}

			log.Logger.Debugf("Set Player to active state")

			l.lock.Lock()
			l.GameStarted = true
			l.lock.Unlock()

			log.Logger.Debugf("Starting game")

			l.Bank.UpdatePublicPlayerBuyIn(l.PublicPlayers)
			l.round.Start(l.Players, l.PublicPlayers, l.dealer)

			log.Logger.Debugf("Game finished")

			l.lock.Lock()
			l.GameStarted = false
			l.lock.Unlock()

			log.Logger.Debugf("Remove players after round")
			l.RemoveAfterRound()

			//empty the starting channel.
		L:
			for {
				select {
				case <-l.c:
				default:
					break L
				}
			}
		}
	}()
}
