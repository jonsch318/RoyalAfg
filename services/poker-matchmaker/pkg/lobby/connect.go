package lobby

import (
	"context"
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/serviceconfig"
)

func (m *Manager) Connect(id string) (*TicketRequestResult, error) {
	list, err := m.agonesClient.AgonesV1().GameServers("default").List(context.Background(), metav1.ListOptions{
		TypeMeta: metav1.TypeMeta{
			Kind:       "agones.dev/v1",
			APIVersion: "GameServer",
		},
		LabelSelector:       fmt.Sprintf("game=poker,lobbyId=%v", id),
		Watch:               false,
		AllowWatchBookmarks: false,
	})
	if err != nil {
		m.logger.Errorw("err during connect", "error", err)
		return nil, err
	}
	m.logger.Infow("Agones list results", "id", id, "list", list)

	if len(list.Items) < 1 {
		return nil, fmt.Errorf("no lobbies found with the given id")
	}

	addresses := viper.GetStringSlice(serviceconfig.NodeIPAddresses)

	adr := list.Items[0].Status.Address
	for _, address := range addresses {
		if err2 := m.PingHealth(fmt.Sprintf("%s:%v", address, list.Items[0].Status.Ports[0].Port)); err2 == nil {
			log.Logger.Debugf("Poker Address found of addresses %v => %v", addresses, address)
			adr = address
			break
		}
		log.Logger.Warnf("Poker Address was not valid %v => %v", addresses, address)
	}

	return &TicketRequestResult{Address: fmt.Sprintf("%s:%v", adr, list.Items[0].Status.Ports[0].Port), LobbyId: id}, nil
}

type HealthPingResponse struct {
	Count   int    `json:"count"`
	Running bool   `json:"running"`
	LobbyID string `json:"lobbyId"`
}

func (m *Manager) PingHealth(addr string) error {
	res, err := http.Get(fmt.Sprintf("http://%v/api/poker/health", addr))

	if err != nil {
		return err
	}

	js, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		return err
	}

	var ping HealthPingResponse
	err = json.Unmarshal(js, &ping)
	if err != nil {
		return err
	}

	return nil
}
