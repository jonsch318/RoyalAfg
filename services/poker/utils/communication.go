package utils

import (
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"time"
)

//SendToAll is a utitlity for sending an event (message) to an entire array (lobby) of players.
func SendToAll(players []models.Player, event *models.Event) {
	for i := range players {
		players[i].Out <- event.ToRaw()
	}
}

//SendToPlayerInList is a utility for sending an event (message) to a specific player in a slice
func SendToPlayerInList(players []models.Player, i int, event *models.Event) {
	players[i].Out <- event.ToRaw()
}

//SendToPlayer is a utility for sending an event (message) to a given player
func SendToPlayer(player *models.Player, event *models.Event) {
	player.Out <- event.ToRaw()
}

func SendToChanTimeout(out chan []byte, event *models.Event) error {
	return SendToChanTimeoutD(out, event, 1*time.Minute)
}

func SendToChanTimeoutD(out chan []byte, event *models.Event, d time.Duration) error {
	timer := time.NewTimer(d)
	select {
	case out <- event.ToRaw():
		return nil
	case <-timer.C:
		return errors.New("Send Timeout")
	}
}

//WaitUntilEvent waits for a player event before a timeout succeedes.
func WaitUntilEvent(player *models.Player) (*models.Event, error) {
	return WaitUntilEventD(player.In, 1*time.Minute)
}

//WaitUntilEventD waits for a player event before a given timeout duration succeedes.
func WaitUntilEventD(in chan *models.Event, d time.Duration) (*models.Event, error) {
	timer := time.NewTimer(d)
	select {
	case e := <-in:
		timer.Stop()
		return e, nil
	case <-timer.C:
		timer.Stop()
		return nil, errors.New("Action Timeout")
	}
}
