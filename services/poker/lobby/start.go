package lobby

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"runtime/debug"
	"sync/atomic"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"

	"github.com/spf13/viper"
)

type Once struct {
	done uint32
}

func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	//We should mutex lock here, but because we only set the done at the start and call reset at the end it is fine for this application.
	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1) //change from resync (github.com/matryer/resync)
		f()
	}
}

func (o *Once) CheckIfAlreadyDone() bool {
	return atomic.LoadUint32(&o.done) == 1
}

func (o *Once) Reset() {
	//We should mutex lock here, but because we only set the done at the start and call reset at the end it is fine for this application.
	atomic.StoreUint32(&o.done, 0)
}

//Start  starts a poker game
func (l *Lobby) Start() {
	// SETUP
	defer func() {
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v Stacktrace: \n %s", r, string(debug.Stack()))
		}
	}()

	if l.once.CheckIfAlreadyDone() {
		log.Logger.Warnf("Start once already done")
		return
	}

	l.once.Do(func() {
		defer func() {
			log.Logger.Infof("Reset Once")
			l.once.Reset()
		}()
		for l.Count() >= viper.GetInt(serviceconfig.PlayersRequiredForStart) {
			log.Logger.Debug("Start timer")

			timer := time.NewTimer(time.Duration(viper.GetInt(serviceconfig.GameStartTimeout)) * time.Second)
			<-timer.C

			log.Logger.Debugf("Preparing for game")

			l.RemoveLeftPlayers()

			l.PrepareForRound()

			log.Logger.Debugf("Game setup playercount: %d", len(l.Players))

			if viper.GetBool(serviceconfig.NeedEnterToStart) {
				log.Logger.Warnf("Needs input to start")
				reader := bufio.NewReader(os.Stdin)
				fmt.Printf("ENTER for game start \n")
				_, _, _ = reader.ReadLine()
			}

			if len(l.Players) < viper.GetInt(serviceconfig.PlayersRequiredForStart) || len(l.Players) == 0 {
				// Not enough players to start
				log.Logger.Warnf("Not enough players to continue")
				return
			}

			if l.dealer < 0 {
				rand.Seed(time.Now().UnixNano())
				l.dealer = rand.Intn(len(l.Players))
			} else {
				l.dealer = (l.dealer + 1) % len(l.Players)
			}

			log.Logger.Debugf("Dealer chosen %v", l.dealer)

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
			l.RemoveLeftPlayers()
		}
	})

	log.Logger.Debugf("Start exited. This can mean not enough players for start or the once was already called")
	//Lobby has to not enough players to continue. Returning to a waiting state
	utils.SendToAll(l.Players, events.NewLobbyPauseEvent(l.PublicPlayers, l.Count()))
}
