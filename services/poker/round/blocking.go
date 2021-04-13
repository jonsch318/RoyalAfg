package round

import (
	"sort"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

type BlockingList struct {
	list []int
}

func NewBlockingList(list []int) *BlockingList {
	return &BlockingList{
		list: list,
	}
}

func (l *BlockingList) CheckIfEmpty() bool {
	return len(l.list) == 0
}

func (l *BlockingList) RemoveBlocking(i int) {
	l.list = append(l.list[:i], l.list[i+1:]...)
}

func (l *BlockingList) AddBlocking(k int) error {
	isOn := false
	for _, n := range l.list {
		if n == k {
			isOn = true
		}
	}
	if !isOn {
		l.list = append(l.list, k)
	}
	sort.Slice(l.list, func(i, j int) bool {
		return l.list[i] < l.list[j]
	})

	return nil
}

func (l *BlockingList) AddAllButThisBlocking(players []models.Player, k int, bank bank.Interface) []int {
	l.list = nil
	for j := 1; j <= len(players); j++ {
		i := (j + k) % len(players)
		if players[i].Active && i != k && !bank.IsAllIn(players[i].ID) {
			l.list = append(l.list, i)
		}
	}
	log.Logger.Debugf("After raise blocking: %v", len(l.list))
	return l.list
}

func (l *BlockingList) GetNext(removed bool, i int) int {
	if removed {
		return i % len(l.list)
	} else {
		return (i + 1) % len(l.list)
	}
}

func (l *BlockingList) Get(i int) int {
	return l.list[i]
}

func (l *BlockingList) Length() int {
	return len(l.list)
}
