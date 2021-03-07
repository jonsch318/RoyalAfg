package showdown

import (
	"github.com/stretchr/testify/assert"

	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
	"testing"
)

var result int

func benchmarkRankSpecificHand(hand []models.Card, b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r =rankSpecificHand(hand)
	}
	result = r //To disable the compiler removing the function call
}

func BenchmarkRoyalFlushRank(b *testing.B) {benchmarkRankSpecificHand(getRoyalFlush(),b)}
func BenchmarkRoyalPair(b *testing.B) {benchmarkRankSpecificHand(getLowPair(),b)}
func BenchmarkRoyalQuad(b *testing.B) {benchmarkRankSpecificHand(getQuad(),b)}

func TestRankSpecificHand(t *testing.T) {
	t.Run("TripleVsPair", func(t *testing.T) {
		rank2Pair := rankSpecificHand(getLowPair())
		rank3Pair := rankSpecificHand(getTriple())

		assert.Greater(t, rank3Pair, rank2Pair)
	})

	t.Run("HighHighCardVsLowHighCard", func(t *testing.T) {
		rankLowHighCard := rankSpecificHand(getLowHighCard())
		rankHighHighCard :=  rankSpecificHand(getHighHighCard())
		assert.Greater(t, rankHighHighCard, rankLowHighCard)
	})
	t.Run("FlushVsPair", func(t *testing.T) {
		rank2Pair := rankSpecificHand(getLowPair())
		rankFlush := rankSpecificHand(getFlush())
		assert.Greater(t, rankFlush, rank2Pair)
	})

	t.Run("FullHouseVsPair", func(t *testing.T) {
		rank2Pair := rankSpecificHand(getLowPair())
		rankFullHouse := rankSpecificHand(getFullHouse())
		assert.Greater(t, rankFullHouse, rank2Pair)
	})

	t.Run("FullHouseVsFlush", func(t *testing.T) {
		rankFlush := rankSpecificHand(getFlush())
		rankFullHouse := rankSpecificHand(getFullHouse())
		assert.Greater(t, rankFullHouse, rankFlush)
	})

	t.Run("RoyalFlushVsFullHouse", func(t *testing.T) {
		rankFullHouse := rankSpecificHand(getFullHouse())
		rankRoyalFlush := rankSpecificHand(getRoyalFlush())
		assert.Greater(t, rankRoyalFlush, rankFullHouse)
	})

	t.Run("QuadVsTriple", func(t *testing.T) {
		rankTriple := rankSpecificHand(getTriple())
		rankQuad := rankSpecificHand(getQuad())
		assert.Greater(t, rankQuad, rankTriple)
	})

	t.Run("QuadVsFullHouse", func(t *testing.T) {
		rankFullHouse := rankSpecificHand(getFullHouse())
		rankQuad := rankSpecificHand(getQuad())
		assert.Greater(t, rankQuad, rankFullHouse)
	})

	t.Run("FlushVsStraight", func(t *testing.T) {
		rankStraight := rankSpecificHand(getStraight())
		rankFlush := rankSpecificHand(getFlush())
		assert.Greater(t, rankFlush, rankStraight)
	})
	t.Run("StraightFlushVsQuad", func(t *testing.T) {
		rankQuad := rankSpecificHand(getQuad())
		rankStraightFlush := rankSpecificHand(getStraightFlush())
		assert.Greater(t, rankStraightFlush, rankQuad)
	})
	t.Run("TwoPairVsTwoPair", func(t *testing.T) {
		twoPairHigh := []models.Card{
			c(0,5),
			c(1,5),
			c(0,3),
			c(0,3),
			c(2,8),
		}
		rankTwoPairLow := rankSpecificHand(getTwoPair())
		rankTwoPairHigh := rankSpecificHand(twoPairHigh)
		assert.Greater(t, rankTwoPairHigh, rankTwoPairLow)
	})
}

func getTwoPair() []models.Card {
	return []models.Card{
		c(0,5),
		c(1,5),
		c(0,3),
		c(0,3),
		c(2,2),
	}
}

//Used for lower pairs (33)
func getLowPair() []models.Card {
	return []models.Card{
		c(0,0),
		c(1,5),
		c(0,1),
		c(0,2),
		c(2,1),
	}
}

func getTriple() []models.Card {
	return []models.Card{
		c(0,1),
		c(1,3),
		c(0,3),
		c(2,4),
		c(0,3),
	}
}

func getHighHighCard() []models.Card {
	return []models.Card{
		c(0,1),
		c(1,10),
		c(0,3),
		c(2,4),
		c(0,7),
	}
}

func getLowHighCard() []models.Card {
	return []models.Card{
		c(0,1),
		c(1,4),
		c(0,3),
		c(2,6),
		c(0,0),
	}
}

func getFlush() []models.Card {
	return []models.Card{
		c(0,1),
		c(0,4),
		c(0,3),
		c(0,6),
		c(0,0),
	}
}

func getStraight() []models.Card {
	return []models.Card{
		c(0,5),
		c(1,10),
		c(0,6),
		c(2,9),
		c(3,8),
	}
}

func getFullHouse() []models.Card {
	return []models.Card{
		c(0,8),
		c(1,8),
		c(0,0),
		c(2,8),
		c(3,0),
	}
}


func getQuad() []models.Card {
	return []models.Card{
		c(0,8),
		c(1,8),
		c(0,2),
		c(2,8),
		c(3,8),
	}
}

func getStraightFlush() []models.Card {
	return []models.Card{
		c(1,2),
		c(1,6),
		c(1,3),
		c(1,5),
		c(1,4),
	}
}

func getRoyalFlush() []models.Card{
	return []models.Card{
		c(0,12),
		c(0,8),
		c(0,9),
		c(0,11),
		c(0,10),
	}
}