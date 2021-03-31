package main

import (
	"fmt"
	"strconv"

	coresdk "agones.dev/agones/pkg/sdk"

	pokerModels "github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/lobby"
)

func SetLobby(b *bank.Bank, lobbyInstance *lobby.Lobby, gs *coresdk.GameServer) error {
	labels := gs.GetObjectMeta().GetLabels()
	min, err := GetFromLabels("min-buy-in", labels)
	if err != nil {
		return err
	}

	max, err := GetFromLabels("max-buy-in", labels)
	if err != nil {
		return err
	}

	blind, err := GetFromLabels("blind", labels)
	if err != nil {
		return err
	}

	index, err := GetFromLabels("class-index", labels)
	if err != nil {
		return err
	}

	lobbyID, ok := labels["lobbyId"]
	if !ok {
		return fmt.Errorf("key needed [%v]", "lobbyId")
	}
	b.RegisterLobby(lobbyID)

	lobbyInstance.RegisterLobbyValue(&pokerModels.Class{
		Min:   min,
		Max:   max,
		Blind: blind,
	}, index, lobbyID)
	return nil
}

func GetFromLabels(key string, labels map[string]string) (int, error) {
	valString, ok := labels[key]

	if !ok {
		return 0, fmt.Errorf("key needed [%v]", key)
	}

	val, err := strconv.Atoi(valString)
	if err != nil {
		return 0, err
	}

	return val, nil
}