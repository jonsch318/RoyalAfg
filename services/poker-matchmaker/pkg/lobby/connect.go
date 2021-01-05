package lobby

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"net/http"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

func (m *Manager) Connect(lobby models.Lobby) (string, error) {
	addr, err := m.rdg.Get(context.Background(), lobby.LobbyId).Result()
	if err != nil {
		return "", err
	}

	running := false
	running, err = m.PingHealth(addr)

	if err != nil || !running {
		m.rdg.Del(context.Background(), lobby.LobbyId)
		return "", fmt.Errorf("error during ping: %s", err)
	}

	//GameServer pointed by addr is healthy
	return addr, err
}

type HealthPingResponse struct {
	Count int `json:"count"`
	Running bool `json:"running"`
}

func (m *Manager) PingHealth(addr string) (bool, error){
	res, err := http.Get(fmt.Sprintf("http://%v/api/poker/health", addr))

	if err != nil {
		return false, err
	}

	js, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return false, err
	}


	var ping HealthPingResponse
	err = json.Unmarshal(js, &ping)
	if err != nil {
		return false, err
	}

	return ping.Running, nil
}

