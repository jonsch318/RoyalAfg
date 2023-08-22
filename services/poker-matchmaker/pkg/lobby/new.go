package lobby

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	allocationv1 "agones.dev/agones/pkg/apis/allocation/v1"
	"github.com/spf13/viper"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/poker/models"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/serviceconfig"
)

// NewLobby allocates a new GameServer for a new RoundId
func (m *Manager) NewLobby(classIndex int) (*TicketRequestResult, error) {

	if classIndex < 0 || classIndex >= len(m.classes) {
		return nil, errors.New("the class index has to be a valid class registered at service start")
	}

	id := newID()
	//model := models.NewLobby(id, class)

	gsa := m.agonesClient.AllocationV1()

	class := m.classes[classIndex]
	serverLabels := make(map[string]string)
	serverLabels["lobbyId"] = id
	serverLabels["players"] = "0"
	serverLabels["min-buy-in"] = strconv.Itoa(class.Min)
	serverLabels["max-buy-in"] = strconv.Itoa(class.Max)
	serverLabels["blind"] = strconv.Itoa(class.Blind)
	serverLabels["class-index"] = strconv.Itoa(classIndex)

	alloc := &allocationv1.GameServerAllocation{
		ObjectMeta: v1.ObjectMeta{
			Name:      id,
			Namespace: "default",
		},
		Spec: allocationv1.GameServerAllocationSpec{
			Required: allocationv1.GameServerSelector{
				LabelSelector: v1.LabelSelector{
					MatchLabels: map[string]string{
						"game": "poker",
					},
					MatchExpressions: nil,
				},
			},
			Preferred: nil,
			MetaPatch: allocationv1.MetaPatch{
				Labels:      serverLabels,
				Annotations: nil,
			},
		}}

	allocationResponse, err := gsa.GameServerAllocations("default").Create(context.Background(), alloc, v1.CreateOptions{})

	m.logger.Warnw("Allocation", "error", err, "lobbyId", id)

	if err != nil {
		return nil, err
	}

	if allocationResponse.Status.GameServerName == "" || len(allocationResponse.Status.Ports) <= 0 {
		return nil, errors.New("no new server can be allocated")
	}

	m.lobbies[classIndex] = append(m.lobbies[classIndex], models.LobbyBase{
		LobbyID:     id,
		Class:       &m.classes[classIndex],
		ClassIndex:  classIndex,
		PlayerCount: 0,
	})

	addresses := viper.GetStringSlice(serviceconfig.NodeIPAddresses)
	addr := allocationResponse.Status.Address

	for _, address := range addresses {
		if err2 := m.PingHealth(fmt.Sprintf("%s:%v", address, allocationResponse.Status.Ports[0].Port)); err2 == nil {
			log.Logger.Debugf("Poker Address found of addresses %v => %v", addresses, address)
			addr = address
			break
		}
		log.Logger.Warnf("Poker Address was not valid %v => %v", addresses, address)
	}

	return &TicketRequestResult{
		Address: fmt.Sprintf("%s:%v", addr, allocationResponse.Status.Ports[0].Port),
		LobbyId: id,
	}, nil
}

const idLength = 7
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// newID generates a new ID for a new RoundId. RoundId ID are composed of letters for easy share.
func newID() string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(idLength)
	for i := 0; i < idLength; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
