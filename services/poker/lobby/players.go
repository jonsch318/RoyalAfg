package lobby

import (
	"errors"
	"runtime/debug"

	"github.com/spf13/viper"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/events"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/serviceconfig"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/utils"
)

//Join adds the player to the waiting queue and starts the queue emptying if possible.
func (l *Lobby) Join(player *models.Player) error {

	//Check if lobby is empty, if it is set state to allocated.
	if l.Count() <= 0 {
		//Set Gameserver state to allocated.
		if err := l.sdk.Allocate(); err != nil {
			log.Logger.Errorw("error during allocation", "error", err)
		}
	}

	//Send Lobby Info to player even before joining.
	err := utils.SendToPlayerTimeout(player, events.NewLobbyInfoEvent(
		l.LobbyID,
		l.Count(),
		viper.GetInt(serviceconfig.PlayersRequiredForStart),
		l.Class.Max,
		l.Class.Min,
		l.Class.Blind,
		viper.GetInt(serviceconfig.GameStartTimeout),
		l.GameStarted,
	))
	if err != nil {
		//Could not send to player
		return errors.New("could not send to player")
	}

	//Add player to queue
	l.PlayerQueue.Enqueue(player)
	l.SetPlayerCountLabel()

	l.FillLobbyPosition()

	//Start game if not already started.
	if !l.GetGameStarted() && l.Count() >= viper.GetInt(serviceconfig.PlayersRequiredForStart) {
		log.Logger.Debugf("Calling start")
		go l.Start() //Call start in seperate routine, so that this routine can still add players.
	}
	return nil
}

//FillLobbyPosition recursively adds players to the lobby until the maximum of 10
func (l *Lobby) FillLobbyPosition() {
	if l.PlayerCount >= 10 {
		//No new players can be added
		return
	}

	log.Logger.Debugf("Dequeueing player")
	player := l.PlayerQueue.Dequeue()
	if player == nil {
		//No player in queue exiting
		log.Logger.Debug("dequeued player is nil => Queue empty")
		return
	}
	public := player.ToPublic()

	//Next element will be index last element + 1.
	playerIndex := len(l.Players)

	//Check if a player with the given id already exists.
	registeredIndex := l.FindPlayerByID(player.ID)
	if registeredIndex != -1 {
		//Player with the same index
		register := &l.Players[registeredIndex]
		if register.Left {
			//Player reconnected we switch to the new connection
			register.Close = player.Close
			register.Out = player.Out
			register.In = player.In
			register.Left = false

			err := utils.SendToPlayerInListTimeout(l.Players, registeredIndex, events.NewJoinSuccessEvent(l.PublicPlayers, registeredIndex, public.BuyIn, l.GameStarted))
			if err != nil {
				log.Logger.Infof("Could not send join success event to player. Trying to remove player now.")
				if err := l.RemovePlayerByID(l.Players[playerIndex].ID); err != nil {
					log.Logger.Errorw("Could not remove player after failing to send join success", "id", l.Players[playerIndex].ID, "error", err)
				}
				l.FillLobbyPosition()
				return
			}

			go l.watchPlayerConnClose(registeredIndex, player.ID)

			l.FillLobbyPosition()
			return
		} else {
			log.Logger.Infof("player with id [%s] tried entering twice", player.ID)
			l.FillLobbyPosition()
			return
		}
	}

	err := l.addPlayer(player, public)
	if err != nil {
		log.Logger.Infof("Could not send join success event to player. Trying to remove player now.")
		if err := l.RemovePlayerByID(l.Players[playerIndex].ID); err != nil {
			log.Logger.Errorw("Could not remove player after failing to send join success", "id", l.Players[playerIndex].ID, "error", err)
		}
		l.FillLobbyPosition()
		return
	}

	log.Logger.Debugf("Start watching player close and sending event")

	//Start CloseWatching
	go l.watchPlayerConnClose(playerIndex, player.ID)

	//Update player count label (matchmaker)
	l.SetPlayerCountLabel()

	l.FillLobbyPosition()
}

//addPlayer is a helper function to add a player to the playerlist and bank
func (l *Lobby) addPlayer(player *models.Player, public *models.PublicPlayer) error {

	log.Logger.Debugf("Adding a player to playerlist")
	playerIndex := len(l.Players)
	l.Players = append(l.Players, *player)
	l.PublicPlayers = append(l.PublicPlayers, *public)
	l.PlayerCount++
	log.Logger.Debugf("Adding player to poker bank")

	l.Bank.AddPlayer(player)
	l.Bank.UpdatePublicPlayerBuyIn(l.PublicPlayers)

	public.SetBuyIn(l.Bank.GetPlayerWallet(public.ID))

	//Send to currently active players. The joining player is not included. He will get a different confirmation
	utils.SendToAll(l.Players, events.NewPlayerJoinEvent(public, len(l.Players)-1, len(l.Players), l.GameStarted))
	//Send join confirmation to player
	return utils.SendToPlayerInListTimeout(l.Players, playerIndex, events.NewJoinSuccessEvent(l.PublicPlayers, playerIndex, public.BuyIn, l.GameStarted))
}

//WatchPlayerConnClose watches the close channel and removes the player when leaving.
func (l *Lobby) watchPlayerConnClose(playerIndex int, id string) {
	defer func() {
		//Player removal was and is the most crashed situation of the game.
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v Stacktrace: \n %s", r, string(debug.Stack()))
		}
	}()

	//wait for closing message
	m, ok := <-l.Players[playerIndex].Close
	if !ok {

		//find player after close. (Player array could have changed so playerIndex is out of date)
		i := l.FindPlayerByID(id)
		if i < 0 {
			//player not found. Should not happen
			log.Logger.Errorf("Could not find player after close message. Indicating player close is called twice %v", id)
			return
		}

		log.Logger.Debugf("Close channel closed... Indicating player left")
		log.Logger.Warnf("REMOVING player %v", id)
		//remove from lobby when closing
		err := l.RemovePlayerByID(id)
		if err != nil {
			log.Logger.Errorw("error during removal", "id", id, "error", err)
		}
	} else {
		log.Logger.Warnf("Something was send to close channel. Message: %v", m)
	}
}
