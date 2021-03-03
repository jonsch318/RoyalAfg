package showdown

import (
	"github.com/JohnnyS318/RoyalAfgInGo/services/poker/models"
)

//rankSpecificHand generates a rank identification number for 5 card array out of the 7 cards.
func rankSpecificHand(cards []models.Card) int {

	// 1 Byte Number:
	// 4 MSB: Describe Hand State (0: High Card - 8: Straight Flush)
	// 4 LSB: Describe State Identifier via Highest Card (0: Card Value 2 - 12: Card Value Ace)
	// Hand Identifier ranges from 0b00000000 (Theoretically Highest Card Two) to 0b10011101 (Straight Flush with Ace -> Royal Flush)
	// If a given Hand h1 is better than another h2, h1 > h2 always yields true (minor exceptions apply).
	identifier := cards[0].Value

	startColor := cards[0].Color // Needed for checking flushes (all the same color).
	lowPair1 := -1
	lowPair2 := -1
	overflow := -1

	y := 0 // Y - Method
	firstCard := identifier // Value of the first card

	mostSignificantValue := -1 //relevant for pairs. if set
	secondSignificantValue := -1

	minValue := identifier


	validStates := 0b111111111 // Initially all card states are valid.
	for i := 1; i < 5; i++ {
		currentValue := cards[i].Value //value of

		//Check for flush
		if (cards[i].Color) != startColor {
			validStates = validStates & 0b011011111
		}

		//Check minimum Value
		if currentValue < minValue {
			minValue = currentValue
		}

		//Update Y-State
		if firstCard == currentValue {
			y++
		} else {
			y--
		}



		if currentValue == identifier || currentValue == lowPair1 || currentValue == lowPair2 || currentValue == overflow {

			//Current was previously visited (is already in hand) (pair)
			validStates = validStates & 0b011001110
			if currentValue < mostSignificantValue {
				secondSignificantValue = currentValue
			} else if currentValue > mostSignificantValue {
				secondSignificantValue = mostSignificantValue
				mostSignificantValue = currentValue
			}


		} else {
			//Current was not seen before (not a pair)
			overflow = lowPair2 // If over has a value => 4 different cards
			if currentValue > identifier {
				//Move all value state vars
				lowPair2 = lowPair1
				lowPair1 = identifier
				identifier = currentValue
			} else if currentValue > lowPair1 {
				//Move lowPair1 to lowPair2
				lowPair2 = lowPair1
				lowPair1 = currentValue
			} else if currentValue > lowPair2 {
				//Nothing to move
				lowPair2 = currentValue
			} else {
				//Current is smaller is previous cards.
				overflow = currentValue
			}

			//Check if we have 4 or 5 different cards.
			if overflow != -1 {
				validStates = validStates & 0b100100011
			}
		}
	}

	//____further exclusion logic_____
	//This algorithm works with excluding all impossible cases

	//2 different cards
	if lowPair2 == -1 {
		//only 2 different cards... full house or 4 pair
		validStates &= 0b011000000
	}else if overflow == -1 {
		//only 3 different cards => only 2 pair and 3 pair
		validStates &= 0b000001100
	}

	if y == 0 || y == -2 {
		//no full house
		validStates &= 0b101111111
	}else if y == 2 || y == 4{
		//no 4 pair
		validStates &= 0b110111111
	}

	//secondSignificantValue has a value when 4 pair
	if secondSignificantValue == -1 {
		//No 3 pair
		validStates = validStates & 0b111111011
	} else {
		//No 4 pair
		validStates = validStates & 0b111110111
	}

	//Check for range between min value and max value is 4 => (e.g. 2-6)
	if mostSignificantValue == -1 && identifier-minValue == 4 {
		//mostSignificant is high pair.
		validStates = validStates & 0b100100000
	} else if mostSignificantValue == -1 {
		validStates = validStates & 0b000100001
	}

	//correctedState is the valid state from 0-8 0 is lowest; 8 is royal flush.
	correctedState := 0
	for correctedState = 8; correctedState >= 0; correctedState-- {
		if (validStates & (1 << correctedState)) != 0 {
			break
		}
	}

	if mostSignificantValue == -1 {
		mostSignificantValue = identifier
	}
	if secondSignificantValue == -1 {
		if mostSignificantValue == -1 {
			secondSignificantValue = lowPair1
		} else {
			secondSignificantValue = lowPair2
		}
	}

	//log.Printf("Rank Cards: %v => %v", cards, (correctedState<<8)+(mostSignificantValue<<4)+secondSignificantValue)

	return (correctedState << 8) + (mostSignificantValue << 4) + secondSignificantValue

}