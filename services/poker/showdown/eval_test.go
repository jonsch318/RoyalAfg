package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"testing"
)

func TestEvaluate(t *testing.T) {
	//setup
	players := []models.Player{{ID: "a", Active: true}, {ID: "b", Active: true}}
	cards := make(map[string][2]models.Card)
	cards["a"] = [2]models.Card{{Color: 2, Value: 11}, {Color: 1, Value: 11}}
	cards["b"] = [2]models.Card{{Color: 3, Value: 8}, {Color: 0, Value: 1}}
	board := [5]models.Card{{Color: 0, Value: 11}, c(3, 4), c(1, 4), c(0, 7), c(2, 2)}

	winners := Evaluate(players, cards, board)

	if len(winners) > 1 || len(winners) < 1 {
		t.Errorf("The winners do not match %v", winners)
		return
	}

	if winners[0] != "a" {
		t.Errorf("The winners do not match [%v] != %v", "a", winners[0])
		return
	}
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
