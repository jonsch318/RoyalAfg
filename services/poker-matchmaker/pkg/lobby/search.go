package lobby


import (
	"errors"
	"sort"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker-matchmaker/pkg/models"
)

func (m *Manager) SearchWithClass(class int) ([]models.Lobby, error) {

	if m.classes == nil {
		return nil, errors.New("no registered buy in classes")
	}

	selection := m.lobbies[class]

	//Copy selection
	rank := make([]models.Lobby, len(selection))
	copy(rank, selection)

	//Sort less [2,3,4,5, etc...]
	sort.SliceStable(rank, func(i, j int) bool {
		return biasForX(rank[i].Players, 9) < biasForX(rank[i].Players, 9)
	})

	//ordered list of lobbies to try
	return rank, nil
}

//
func biasForX(i, x int) int {
	if i > x {
		return 2*x - i
	}
	return i
}

func (m *Manager) SearchWithParams(min, max, blind int) ([]models.Lobby, error){

	if m.classes == nil {
		return nil, errors.New("no registered buy in classes")
	}

	found := make([]int, 0)
	for i,v := range m.classes {
		if v.Blind == blind && min >= v.Min {
			found = append(found, i)
		}
	}

	smallest :=
		struct {
			i int
			m int
		}{0,0}
	for j := 0; j < len(found); j++ {
		c :=  m.classes[found[j]].Max
		if smallest.m >= c {
			smallest =  struct {
				i int
				m int
			}{j, c}
		}
	}

	return m.SearchWithClass(smallest.i)
}