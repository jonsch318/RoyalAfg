package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/jonsch318/royalafg/pkg/log"
	"github.com/jonsch318/royalafg/services/poker/models"
)

// SendToAll is a utility for sending an event (message) to an entire array (lobby) of players.
func SendToAll(players []models.Player, event *models.Event) {
	for i := range players {
		err := SendToPlayerInListTimeout(players, i, event)
		if err != nil {
			log.Logger.Warnf("Player [%v] already left. Err: %v", players[i].ID, err)
		}
	}
}

// SendToPlayerInList is a utility for sending an event (message) to a specific player in a slice
func SendToPlayerInListTimeout(players []models.Player, i int, event *models.Event) error {
	if !players[i].Left {
		return SendToChanTimeout(players[i].Out, event)
	}
	return fmt.Errorf("player [%v] already left", players[i].ID)
}

// SendToPlayer is a utility for sending an event (message) to a given player
func SendToPlayerTimeout(player *models.Player, event *models.Event) error {
	if !player.Left {
		return SendToChanTimeout(player.Out, event)
	}
	return errors.New("player already left")
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
		return errors.New("send timeout")
	}
}

// WaitUntilEvent waits for a player event before a timeout succeeds.
func WaitUntilCloseOrEvent(player *models.Player) (*models.Event, error) {
	if !player.Left {
		return WaitUntilCloseOrEventD(player.In, player.Close, 1*time.Minute)
	}
	return nil, errors.New("player has already left")
}

// WaitUntilEventD waits for a player event before a given timeout duration succeeds.
func WaitUntilCloseOrEventD(in chan *models.Event, close chan bool, d time.Duration) (*models.Event, error) {
	if in == nil {
		return nil, errors.New("no events can be transmitted")
	}
	timer := time.NewTimer(d)
	select {
	case e := <-in:
		timer.Stop()
		return e, nil
	case x, ok := <-close:
		log.Logger.Infof("Player left... %v, %v", x, ok)
		timer.Stop()
		return nil, errors.New("player closed connection")
	case <-timer.C:
		timer.Stop()
		return nil, errors.New("action timeout")
	}
}

// WaitUntilEventD waits for a player event before a given timeout duration succeeds.
func WaitUntilEventD(in chan *models.Event, d time.Duration) (*models.Event, error) {
	timer := time.NewTimer(d)
	select {
	case e := <-in:
		timer.Stop()
		return e, nil
	case <-timer.C:
		timer.Stop()
		return nil, errors.New("action timeout")
	}
}
