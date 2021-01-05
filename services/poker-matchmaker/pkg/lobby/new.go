package lobby

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	allocationv1 "agones.dev/agones/pkg/apis/allocation/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	//"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

func (m *Manager) NewLobby(class int) (string, error) {
	id := newId()
	//model := models.NewLobby(id, class)

	gsa := m.agonesClient.AllocationV1()


	alloc := &allocationv1.GameServerAllocation{
		ObjectMeta: v1.ObjectMeta{
			Name:                       id,
			Namespace:                  "royalafg-poker",
		},
		Spec:       allocationv1.GameServerAllocationSpec{
			MultiClusterSetting: allocationv1.MultiClusterSetting{
				Enabled:        false,
				PolicySelector: v1.LabelSelector{
					MatchLabels:      nil,
					MatchExpressions: nil,
				},
			},
			Required:            v1.LabelSelector{
				MatchLabels:      nil,
				MatchExpressions: nil,
			},
			Preferred:           nil,
			MetaPatch:           allocationv1.MetaPatch{
				Labels:      nil,
				Annotations: nil,
			},
		},
	}

	allocationResponse, err := gsa.GameServerAllocations("royalafg-poker").Create(alloc)

	if err != nil {
		return "", err
	}

	addr := allocationResponse.Status.Address
	port := allocationResponse.Status.Ports[0].Port
	return fmt.Sprintf("%s:%v", addr, port), nil
}

const idLength = 7
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func newId() string {
	rand.Seed(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(idLength)
	for i := 0; i < idLength; i++ {
		sb.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return sb.String()
}