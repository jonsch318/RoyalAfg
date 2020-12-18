package round

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/bank"
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"log"
	"sort"
)

func checkIfEmpty(blocking []int) bool {
	return len(blocking) <= 0
}

func removeBlocking(blocking []int, i int) []int {
	b := append(blocking[:i], blocking[i+1:]...)
	return b
}

func addBlocking(blocking []int, k int) error {
	isOn := false
	for _, n := range blocking {
		if n == k {
			isOn = true
		}
	}
	if !isOn {
		blocking = append(blocking, k)
	}
	sort.Slice(blocking, func(i, j int) bool {
		return blocking[i] < blocking[j]
	})

	return nil
}

func addAllButThisBlockgin(blocking []int, players []models.Player, k int, bank *bank.Bank) []int {
	blocking = nil
	for j := 1; j <= len(players); j++ {
		i := (j + k) % len(players)
		if players[i].Active && i != k && !bank.IsAllIn(players[i].ID) {
			blocking = append(blocking, i)
		}
	}
	log.Printf("After raise blocking: %v", blocking)
	return blocking
}
