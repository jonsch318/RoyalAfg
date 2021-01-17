package lobby

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"net/http"
)

func (m *Manager) Connect(id string) (*TicketRequestResult, error) {

	addr, err := m.rdg.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}

	running := false
	running, err = m.PingHealth(addr)

	if err != nil || !running {
		m.rdg.Del(context.Background(), id)
		return nil, fmt.Errorf("error during ping: %s", err)
	}

	//GameServer pointed by addr is healthy
	return &TicketRequestResult{Address: addr, LobbyId: id}, err
}

type HealthPingResponse struct {
	Count   int  `json:"count"`
	Running bool `json:"running"`
}

func (m *Manager) PingHealth(addr string) (bool, error) {
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
