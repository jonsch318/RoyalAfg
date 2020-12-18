package lobbies

import (
	"sort"
)

type ByMin [][2]int

func (b ByMin) Len() int           { return len(b) }
func (b ByMin) Less(i, j int) bool { return b[i][0] < b[j][0] }
func (b ByMin) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func GetBuyInClass(classes [][2]int, buyIn int) (result int) {
	//mid := len(classes) / 2
	i := LinearSearch(classes, buyIn)
	if classes[i][0] > buyIn || classes[i][1] < buyIn {
		return -1
	}
	return i
}

func LinearSearch(classes [][2]int, buyIn int) int {
	for i := 0; i < len(classes); i++ {
		if i < len(classes)-1 {
			if classes[i][0] < buyIn && classes[i+1][0] < buyIn {
				continue
			}
		}
		return i
	}
	return -1
}

func OrderBuyInClasses(classes [][2]int) [][2]int {
	sort.Sort(ByMin(classes))
	return classes
}

//CheckBuyInClass checks the given buyin for the specified class. If the buyin is either too low or too high, false is returned.
func CheckBuyInClass(classes [][]int, class int, buyIn int) bool {
	if buyIn < classes[class][0] || buyIn > classes[class][1] {
		return false
	}
	return true
}
