package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

//rankSpecificHand generates a rank identificiation number for 5 card array out of the 7 cards.
func rankSpecificHand(cards []models.Card) int {

	//Hand identifier explanation:
	// 1 Byte Number:
	// 4 MSB: Describe Hand State (0: High Card - 8: Straight Flush)
	// 4 LSB: Describe State Identifier via Highest Card (0: Card Value 2 - 12: Card Value Ace)
	// Hand Identifier ranges from 0b00000000 (Theoretically Highest Card Two) to 0b10011101 (Straight Flush with Ace -> Royal Flush)
	// If a given Hand h1 is better than another h2, h1 > h2 always yields true.
	identifier := normalizeAce(cards[0].Value)

	sCol := normalizeAce(cards[0].Color)
	lP1 := -1
	lP2 := -1
	over := -1

	y := 0
	org := identifier

	m1 := -1
	m2 := -1

	min := identifier

	validStates := 0b111111111
	for i := 1; i < 5; i++ {
		open := normalizeAce(cards[i].Value)

		if normalizeAce(cards[i].Color) != sCol {
			validStates = validStates & 0b011011111
		}

		if org == open {
			y++
		} else {
			y += -1
		}

		if open == identifier || open == lP1 || open == lP2 || open == over {
			validStates = validStates & 0b011001110
			if m1 == -1 {
				m1 = open
			} else if open < m1 {
				m2 = open
			} else if open > m1 {
				m2 = m1
				m1 = open
			}
		} else {
			over = lP2
			if open > identifier {
				lP2 = lP1
				lP1 = identifier
				identifier = open
			} else if open > lP1 {
				lP2 = lP1
			} else if open > lP1 {
				lP2 = lP1
				lP1 = open
			} else if open > lP2 {
				lP2 = open
			} else {
				over = open
			}

			if over != -1 {
				validStates = validStates & 0b100100011
			}
		}

	}

	if lP2 != -1 {
		validStates = validStates & 0b101111111
	} else if y == 2 || y == -4 {
		validStates = validStates & 0b110111111
	}

	if m2 == -1 {
		validStates = validStates & 0b111111011
	} else {
		validStates = validStates & 0b111110111
	}

	if m1 == -1 && identifier-min == 4 {
		validStates = validStates & 0b100100000
	} else if m1 == -1 {
		validStates = validStates & 0b000100001
	}

	f := 0
	for f = 8; f >= 0; f-- {
		if (validStates & (1 << f)) != 0 {
			break
		}
	}

	if m1 == -1 {
		m1 = identifier
	}
	if m2 == -1 {
		if m1 == -1 {
			m2 = lP1
		} else {
			m2 = identifier
		}
	}

	//log.Printf("Rank Cards: %v => %v", cards, (f<<8)+(m1<<4)+m2)

	return (f << 8) + (m1 << 4) + m2

}
