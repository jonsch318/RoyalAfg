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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

//NewLobby allocates a new GameServer for a new Lobby
func (m *Manager) NewLobby(classIndex int) (string, error) {

	if classIndex < 0 || classIndex >= len(m.classes) {
		return "", errors.New("the class index has to be a valid class registered at service start")
	}

	id := newID()
	//model := models.NewLobby(id, class)

	gsa := m.agonesClient.AllocationV1()

	class := m.classes[classIndex]
	serverLabels := make(map[string]string, 3)
	serverLabels["lobbyId"] = id
	serverLabels["min-buy-in"] = strconv.Itoa(class.Min)
	serverLabels["max-buy-in"] = strconv.Itoa(class.Max)
	serverLabels["Blid"] = strconv.Itoa(class.Blind)

	alloc := &allocationv1.GameServerAllocation{
		ObjectMeta: v1.ObjectMeta{
			Name:      id,
			Namespace: "royalafg-poker",
		},
		Spec: allocationv1.GameServerAllocationSpec{
			Required: v1.LabelSelector{
				MatchLabels: map[string]string{
					"game": "royalafg-poker",
				},
				MatchExpressions: nil,
			},
			Preferred: nil,
			MetaPatch: allocationv1.MetaPatch{
				Labels:      serverLabels,
				Annotations: nil,
			},
		},
	}

	allocationResponse, err := gsa.GameServerAllocations("royalafg-poker").Create(alloc)

	if err != nil {
		return "", err
	}

	ip := allocationResponse.Status.Address
	port := allocationResponse.Status.Ports[0].Port
	addr := fmt.Sprintf("%s:%v", ip, port)

	err = m.rdg.Set(context.Background(), id, addr, 0).Err()

	if err != nil {
		return "", err
	}

	return addr, nil
}

const idLength = 7
const letterBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

//newID generates a new ID for a new Lobby. Lobby ID are composed of letters for easy share.
func newID() string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(idLength)
	for i := 0; i < idLength; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}
