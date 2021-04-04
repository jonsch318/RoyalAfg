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

	log.Logger.Debugf("Gamestart if [%d] > %d && %v", l.Count(), viper.GetInt(serviceconfig.PlayersRequiredForStart), !l.GetGameStarted())
	//Start game if not already started.
	if l.Count() >= viper.GetInt(serviceconfig.PlayersRequiredForStart) && !l.GetGameStarted() {
		log.Logger.Debugf("Calling start")
		l.Start()
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

	//Check if player is unique when required
	if viper.GetBool(serviceconfig.GameRequiresUniquePlayers) && l.FindPlayerByID(player.ID) != 1 {
		log.Logger.Infof("player with id [%s] tried entering twice with the unique player requirement", player.ID)
		//Moving on to next player in queue.
		l.FillLobbyPosition()
	}

	log.Logger.Info("Adding player to player list")
	playerIndex := len(l.Players)
	l.Players = append(l.Players, *player)
	l.PublicPlayers = append(l.PublicPlayers, *public)
	l.PlayerCount++
	log.Logger.Debugf("Adding players internal bank entry")

	l.Bank.AddPlayer(player)
	l.Bank.UpdatePublicPlayerBuyIn(l.PublicPlayers)

	public.SetBuyIn(l.Bank.GetPlayerWallet(public.ID))

	//Send to currently active players. The joining player is not included. He will get a different confirmation
	utils.SendToAll(l.Players, events.NewPlayerJoinEvent(public, len(l.Players)-1, len(l.Players), l.GameStarted))
	//Send join confirmation to player
	err := utils.SendToPlayerInListTimeout(l.Players, playerIndex, events.NewJoinSuccessEvent(l.PublicPlayers, playerIndex, public.BuyIn, l.GameStarted))
	if err != nil {
		log.Logger.Infof("Could not send to player")
		if err := l.RemovePlayerByID(l.Players[playerIndex].ID); err != nil {
			log.Logger.Errorw("error during removal", "id", l.Players[playerIndex].ID, "error", err)
		}
		l.FillLobbyPosition()
		return
	}
	log.Logger.Debugf("send event")

	log.Logger.Debugf("started watching player close and sending event")

	//Start CloseWatching
	go l.WatchPlayerConnClose(playerIndex, player.ID)

	//Update player count label (matchmaker)
	l.SetPlayerCountLabel()

	l.FillLobbyPosition()
}

//WatchPlayerConnClose watches the close channel and removes the player when leaving.
func (l *Lobby) WatchPlayerConnClose(playerIndex int, id string) {
	defer func() {
		//Player removal was and is the most crashed situation of the game.
		if r := recover(); r != nil {
			log.Logger.Debugf("recovering in round start from %v Stacktrace: \n %s", r, string(debug.Stack()))
		}
	}()

	//wait for closing message
	_, ok := <-l.Players[playerIndex].Close

	//find player after close. (Player array could have changed so playerIndex is out of date)
	i := l.FindPlayerByID(id)
	if i < 0 {
		//player not found. Should not happen
		log.Logger.Errorw("Could not find player after close message")
		return
	}

	if !ok {
		log.Logger.Warnf("Close channel closed... Indicating player left")
		log.Logger.Infof("Removing player %v", l.Players[i].ID)
		//remove from lobby when closing
		err := l.RemovePlayerByID(l.Players[i].ID)

		if err != nil {
			log.Logger.Errorw("error during removal", "id", l.Players[i].ID, "error", err)
		}
	} else {
		log.Logger.Warnf("Something was send to close channel")
	}
}

//PlayerRemoval removes all players in the removal queue.
func (l *Lobby) PlayerRemoval() {

	player := l.RemovalQueue.Dequeue()
	if player == nil {
		//No player in queue
		return
	}

	if player.ID == "" {
		log.Logger.Debugf("player id nil")
	}
	//Get index of player
	i := l.FindPlayerByID(player.ID)
	if i < 0 {
		log.Logger.Warnw("Id [%v] not in lobby", player.ID)
		return
	}
	public := l.PublicPlayers[i]

	//Remove player from list, public list and bank
	l.Players = append(l.Players[:i], l.Players[i+1:]...)
	l.PublicPlayers = append(l.PublicPlayers[:i], l.PublicPlayers[i+1:]...)
	l.PlayerCount--

	err := l.Bank.RemovePlayer(player.ID)
	if err != nil {
		log.Logger.Errorw("error during removing player from bank", "error", err)
	}

	//Send leave event
	utils.SendToAll(l.Players, events.NewPlayerLeavesEvent(&public, i, len(l.Players), l.GameStarted))

	l.PlayerRemoval()
}
