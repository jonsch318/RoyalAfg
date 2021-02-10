package lobby

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"

	"github.com/spf13/viper"
)

//Start  starts a poker game
func (l *Lobby) Start() {

	// SETUP
	go func() {
		for len(l.Players) >= viper.GetInt(serviceconfig.PlayersRequiredForStart) {

			// Protection against multiple games using a buffered channel
			timer := time.NewTimer(15 * time.Second)
			select {
			case <-l.c:
				// game has already been called this instance is unneccesairy
				return
			case <-timer.C:
			}

			// channel is empty, so the buffer is free to be filled.
			l.c <- true

			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("ENTER for game start \n")
			reader.ReadLine()

			if len(l.Players) < viper.GetInt(serviceconfig.PlayersRequiredForStart) {
				// Not enough players to start
				return
			}

			if l.dealer < 0 {
				rand.Seed(time.Now().UnixNano())
				l.dealer = rand.Intn(len(l.Players))
			} else {
				l.dealer = (l.dealer + 1) % len(l.Players)
			}

			for i := range l.Players {
				// Set player states to active
				l.Players[i].Active = true
			}

			l.lock.Lock()
			l.GameStarted = true
			l.lock.Unlock()

			l.Bank.UpdatePublicPlayerBuyIn(l.PublicPlayers)
			l.round.Start(l.Players, l.PublicPlayers, l.dealer)

			l.lock.Lock()
			l.GameStarted = false
			l.lock.Unlock()

			if l.HasToBeRemoved() {
				l.RemoveAfterGame()
			}
			if l.HasToBeAdded() {
				l.EmptyToBeAdded()
			}
		}
	}()
}
