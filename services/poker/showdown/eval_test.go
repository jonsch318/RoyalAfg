package showdown

import (
	"github.com/stretchr/testify/assert"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"testing"
)

var winnerResult []WinnerInfo
func TestEvaluate(t *testing.T) {
	//setup
	players := []models.Player{{ID: "a", Active: true}, {ID: "b", Active: true}}
	cards := make(map[string][2]models.Card)
	cards["a"] = [2]models.Card{{Color: 2, Value: 11}, {Color: 1, Value: 11}}
	cards["b"] = [2]models.Card{{Color: 3, Value: 8}, {Color: 0, Value: 1}}
	board := [5]models.Card{{Color: 0, Value: 11}, c(3, 4), c(1, 4), c(0, 7), c(2, 2)}

	winners := Evaluate(players, cards, board, 2)

	assert.Len(t, winners, 1)
	assert.Equal(t, "a", winners[0].Player.ID)
}

func BenchmarkEvaluate(b *testing.B) {
	players := []models.Player{{ID: "a", Active: true}, {ID: "b", Active: true}}
	cards := make(map[string][2]models.Card)
	cards["a"] = [2]models.Card{{Color: 2, Value: 11}, {Color: 1, Value: 11}}
	cards["b"] = [2]models.Card{{Color: 3, Value: 8}, {Color: 0, Value: 1}}
	board := [5]models.Card{{Color: 0, Value: 11}, c(3, 4), c(1, 4), c(0, 7), c(2, 2)}
	var winners []WinnerInfo
	for i := 0; i < b.N; i++ {
		winners = Evaluate(players, cards, board, 2)
	}
	winnerResult = winners
}

func TestEvaluatePlayer(t *testing.T) {
	//setup
	cards := make(map[string][2]models.Card)
	cards["a"] = [2]models.Card{c(0, 12), c(0, 10)}
	//cards["b"] = [2]models.Card{c(0, 0), c(1, 11)}
	board := [5]models.Card{c(0, 11), c(0, 8), c(1, 8), c(0, 9), c(0, 5)}

	for _, v := range cards {
		rank := evaluatePlayer(append(v[:], board[:]...))
		t.Logf("Rank player %v", rank)
	}
}

func c(a, v int) models.Card {
	return models.Card{Color: a, Value: v}
}
